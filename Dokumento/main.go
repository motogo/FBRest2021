package main

import (
	"github.com/gorilla/mux"
	"fbrest/Base/apis"
	"log"
	"net/http"
)

func main(){

	router := mux.NewRouter()

	router.HandleFunc("/dokumento/standort/show",apis.GetAllLocations).Methods("GET") 
	router.HandleFunc("/dokumento/anforderungen",apis.GetAllAnforderungen).Methods("GET")
	router.HandleFunc("/db/sql",apis.GetSQL).Methods("GET")
	router.HandleFunc("/db/sqlrows",apis.GetSQLRows).Methods("GET")
	router.HandleFunc("/db/test",apis.TestDBOpenClose).Methods("GET")
	router.HandleFunc("/api/test",apis.TestResponse).Methods("GET")
	router.HandleFunc("/dokumento/standort/insert",apis.InsertLocation).Methods("POST")
	router.HandleFunc("/dokumento/standort/delete",apis.DeleteLocation).Methods("DELETE")


	err := http.ListenAndServe(":4500",router)

	if err != nil {
		log.Println(err)
	}

}