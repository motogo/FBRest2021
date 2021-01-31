package apis

import (
	"fbrest/Base/config"
	"fbrest/Base/models"	
	"fbrest/Dokumento/models"	
	"log"
	"mux"
	
	
    
	_struct "fbrest/Base/struct"
	
	"encoding/json"
	"net/http"
)

func GetAllLocations(response http.ResponseWriter, r *http.Request) {

    log.Println("Apis.GetAllLocations") 
	db, err := config.Conn()
	var Response _struct.ResponseData
	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB:db}
		IsiData, err2 := _models.GetDataTableStandort()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)

		} else {
			Response.Status = http.StatusOK
			Response.Message = "Erfolgreich"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)

		}
	}

}







func GetAllAnforderungen(response http.ResponseWriter, r *http.Request) {

	db, err := config.Conn()
	var Response _struct.ResponseData
	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB:db}
		IsiData, err2 := _models.GetDataTableAnforderungen()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)

		} else {
			Response.Status = http.StatusOK
			Response.Message = "Erfolgreich"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)

		}
	}

}


func DeleteLocation(response http.ResponseWriter, r *http.Request) {
    log.Println("DeleteLocation") 
	setupResponse(&response, r)
	vars := mux.Vars(r)
	log.Println(vars) 
	_id := vars["id"]
	log.Println(_id) 
	db, err := config.Conn()
	var _response _struct.ResponseData
	if err != nil {
		log.Println(err.Error())
	}

	_model := models.Db_init{DB: db}

	resultData := _model.DeleteLocation(_id)

	if resultData == "gagal" {
		_response.Status = http.StatusInternalServerError
		_response.Message = "Gagal Delete Data"
		restponWithJson(response, int(_response.Status), _response)
	} else {
		_response.Status = http.StatusOK
		_response.Message = "Sukses Delete Data"
		restponWithJson(response, int(_response.Status), _response)
	}
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func InsertLocation(response http.ResponseWriter, r *http.Request) {

    
	setupResponse(&response, r)
	var entitiesData _struct.Location
	
	err := json.NewDecoder(r.Body).Decode(&entitiesData)
	
	if err != nil {
		log.Println(err.Error()) 
	}
	
	db, err := config.Conn()
	var Response _struct.ResponseData
	
	
	log.Println("Datax") 
	log.Println(entitiesData.ID) 
	log.Println(entitiesData.BEZ) 
	log.Println(entitiesData.GUELTIG) 
	
	if err != nil {
		log.Println(err.Error()) 
	}
	
	_model := models.Db_init{DB: db}
    
	resultData := _model.InsertLocation(&entitiesData)

	if resultData != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = "Gagal Insert Data"
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		Response.Status = http.StatusOK
		Response.Message = "Succses Insert Data"
		Response.Data = entitiesData
		restponWithJson(response, http.StatusOK, Response)
	}

}

func restponWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


