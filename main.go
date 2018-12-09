package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"	// package needed (for registering its drivers with the database/sql package) but never directly referenced in code
)

const (
	host     = "localhost"
	port     = 5432
	user     = "johndibaggio"
	password = "p4ssw0rd"
	dbname   = "sensor_data"
	sslmode  = "disable"	// require
)

var dataset []DataPoint // TODO: replace with db

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to postgres database %s", dbname)


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

