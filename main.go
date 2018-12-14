package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	_ "github.com/lib/pq"	// package needed (for registering its drivers with the database/sql package) but never directly referenced in code
	"github.com/dibaggioj/sensor-api/database"
	"github.com/dibaggioj/sensor-api/database/table"
	"github.com/dibaggioj/sensor-api/models"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "johndibaggio"
	password = "p4ssw0rd"
	dbname   = "sensor_data"
	sslmode  = "disable"	// require
)

var dataset []models.DataPoint // TODO: replace with db

func main() {



	dbConfig := database.Config{
		Host:"localhost",
		Port: 5432,
		User: "johndibaggio",
		Password: "p4ssw0rd",
		Database: "sensor_data",
		SSLMode: "disable"}

	connection, err := database.New(dbConfig)
	if err != nil {
		panic(err)
	}
	connectionPtr := &connection


	measurementsTableConfig := table.DataTableConfig{Connection: connectionPtr, Name: "measurements", CreateSql: table.SQL_TABLE_CREATION_MEASUREMENTS}
	dataSetTableConfig := table.DataTableConfig{Connection: connectionPtr, Name: "dataset", CreateSql: table.SQL_TABLE_CREATION_DATA_SET}

	measurementsTable, err := table.NewDataTable(measurementsTableConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully created table: " + measurementsTable.Name)
	}

	dataSetTable, err := table.NewDataTable(dataSetTableConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully created table: " + dataSetTable.Name)
	}

	testdata := models.DataPoint{ID: 0,
		Timestamp: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Data: &models.SensorData{Temperature: 89.0, Humidity: 0.43}}

	row, err := dataSetTable.InsertDataPoint(testdata)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Created row %d\n", row.ID)
	}


	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	//	host, port, user, password, dbname, sslmode)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Successfully connected to postgres database %s", dbname)


	dataset = append(dataset, models.DataPoint{ID: 0,
		Timestamp: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Data: &models.SensorData{Temperature: 72.0, Humidity: 0.13}})
	dataset = append(dataset, models.DataPoint{ID: 1,
	Timestamp: time.Date(2018, 11, 24, 20, 34, 58, 651387237, time.UTC),
		Data: &models.SensorData{Temperature: 71.0, Humidity: 0.10}})
	dataset = append(dataset, models.DataPoint{ID: 2, Timestamp: time.Now(), Data: &models.SensorData{Temperature: 74.0, Humidity: 0.23}})

	router := mux.NewRouter()
	router.HandleFunc("/sensordata/api/data", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data/{id}", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data", CreateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", UpdateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", DeleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

