package main

import (
"log"
"net/http"
"github.com/gorilla/mux"
	"time"
)

var dataset []DataPoint // TODO: replace with db

func main() {
	dataset = append(dataset, DataPoint{ID: 0,
		Time: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Data: &SensorData{Temperature: 72.0, Humidity: 0.13}})
	dataset = append(dataset, DataPoint{ID: 1,
	Time: time.Date(2018, 11, 24, 20, 34, 58, 651387237, time.UTC),
		Data: &SensorData{Temperature: 71.0, Humidity: 0.10}})
	dataset = append(dataset, DataPoint{ID: 2, Time: time.Now(), Data: &SensorData{Temperature: 74.0, Humidity: 0.23}})

	router := mux.NewRouter()
	router.HandleFunc("/sensordata/api/data", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data/{id}", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data", CreateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", UpdateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", DeleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

