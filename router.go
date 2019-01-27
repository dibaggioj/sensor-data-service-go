package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func route() {
	router := mux.NewRouter()
	router.HandleFunc("/sensordata/api/data", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data/{id}", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data", CreateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", UpdateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", DeleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
