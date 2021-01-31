package dokumentomodels

import (
	_struct "fbrest/Base/struct"
	"fbrest/Base/config"
	"database/sql"
	"log"
	"null"
	
	
)

type ModelGetData struct {
	DB *sql.DB
}

type Db_init struct {
	DB *sql.DB
}


func (model ModelGetData) GetDataTableStandort() (getStruct []_struct.StructData, err error) {
	row, err := model.DB.Query("select ID, BEZ, SCHLUESSEL FROM TSTANDORT")
	if err != nil {
		return  nil,err
	} else {
	    var nullstr null.String
		nullstr = null.NewString("foo", true)
		log.Println(nullstr.ValueOrZero())
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.ID,
				&data.BEZ,
				&data.SCHLUESSEL)
			log.Println("row.Scan")		
			if err2 != nil {
			    log.Println("err2")
				log.Println(err2.Error())
				return nil, err2
			} else {
				_data := _struct.StructData{
					ID:  data.ID,
					BEZ: data.BEZ,
					SCHLUESSEL:  data.SCHLUESSEL,
						
				}
				log.Println("before append")	
				_isiStruct = append(_isiStruct, _data)
				log.Println("after append")	

			}
		}

		return _isiStruct, nil
	}
}



        

func (model Db_init) InsertLocation(OutputData *_struct.Location) (err error) {
	var dataID = OutputData.ID
	var dataBEZ = OutputData.BEZ
	var dataGUELTIG = OutputData.GUELTIG
	var Count_data int

	CheckDataDate, errDataDate := config.Conn()
	CheckDataDate.QueryRow(`select COUNT(*) from TSTANDORT where ID =?`, dataID).Scan(&Count_data)
	if errDataDate != nil {
		/*log.Println(errDataDate)*/
	}
	if Count_data == 0 {
		insertData, errInsertData := model.DB.Exec(`insert into TSTANDORT(ID,BEZ,GUELTIG) VALUES(?,?,?) `, dataID, dataBEZ, dataGUELTIG)
		if errInsertData != nil {
		/*	return errors.New("Server Disconnect")*/
		} else {
			insertData.LastInsertId()
			return nil
		}

	}
	return nil
}



func (model ModelGetData) GetDataTableAnforderungen() (getStruct []_struct.StructVANFORDERUNGALLEData, err error) {
	row, err := model.DB.Query("select ID, BEZ, BESCHREIBUNG FROM VDOKUMENTALLE")
	if err != nil {
		return  nil,err
	} else {
		var _isiStruct []_struct.StructVANFORDERUNGALLEData
		var data _struct.StructVANFORDERUNGALLEData
		for row.Next() {
			err2 := row.Scan(
				&data.ID,
				&data.BEZ,
				&data.BESCHREIBUNG)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructVANFORDERUNGALLEData{
					ID:  data.ID,
					BEZ: data.BEZ,
					BESCHREIBUNG:  data.BESCHREIBUNG,
				}
				_isiStruct = append(_isiStruct, _data)

			}
		}

		return _isiStruct, nil
	}
}

func (model Db_init) DeleteLocation(id string) string {
	var responseData string
	insertData, errInsertData := model.DB.Exec(`delete from tstandorte where id = ? `, id)
	if errInsertData != nil {
		responseData = "False"
	} else {
		insertData.RowsAffected()
		responseData = "Erfolgreich"
	}
	return responseData
}
