package main

import (
	_ "github.com/lib/pq"	// package needed (for registering its drivers with the database/sql package) but never directly referenced in code
	"github.com/dibaggioj/sensor-api/models"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tkanos/gonfig"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
)

var dataset []models.DataPoint // TODO: replace with db
var db *gorm.DB

func main() {

	var err error

	config := models.Configuration{}
	err = gonfig.GetConf("config/config.development.json", &config)
	if err != nil {
		fmt.Println("Unable to get configuration, exiting...")
		panic(err)
		return
	} else {
		fmt.Printf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s\n", config.Db_Host, config.Db_Port,
			config.Db_User, config.Db_Dbname, config.Db_Password, config.Db_Sslmode)
	}

	db, err = gorm.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.Db_User, config.Db_Password, config.Db_Dbname, config.Db_Host, config.Db_Port, config.Db_Sslmode))
	defer db.Close()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully connected to postgres database %s\n", config.Db_Dbname)
	}

	if !db.HasTable(&models.DataPoint{}) {
		db.CreateTable(&models.DataPoint{})
	}
	//if !db.HasTable("data_points") {
	//    db.CreateTable(&models.DataPoint{})
	//}
	if !db.HasTable(&models.SensorData{}) {
		db.CreateTable(&models.SensorData{})
	}
	//if !db.HasTable("sensor_data") {
	//    db.CreateTable(&models.SensorData{})
	//}


	db.AutoMigrate(&models.DataPoint{}, &models.SensorData{})


	//connection, err := database.New(dbConfig)
	//if err != nil {
	//	panic(err)
	//}
	//connectionPtr := &connection


	//measurementsTableConfig := table.DataTableConfig{Connection: connectionPtr, Name: "measurements", CreateSql: table.SQL_TABLE_CREATION_MEASUREMENTS}
	//dataSetTableConfig := table.DataTableConfig{Connection: connectionPtr, Name: "dataset", CreateSql: table.SQL_TABLE_CREATION_DATA_SET}
	//
	//measurementsTable, err := table.NewDataTable(measurementsTableConfig)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("Successfully created table: " + measurementsTable.Name)
	//}
	//
	//dataSetTable, err := table.NewDataTable(dataSetTableConfig)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("Successfully created table: " + dataSetTable.Name)
	//}
	//
	//testdata := models.DataPoint{ID: 0,
	//	Timestamp: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC),
	//	Data: &models.SensorData{Temperature: 89.0, Humidity: 0.43}}
	//
	//row, err := dataSetTable.InsertDataPoint(testdata)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Printf("Created row %d\n", row.ID)
	//}


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


	//dataset = append(dataset, models.DataPoint{ID: 0,
	//	Timestamp: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC),
	//	Data: &models.SensorData{Temperature: 72.0, Humidity: 0.13}})
	//dataset = append(dataset, models.DataPoint{ID: 1,
	//Timestamp: time.Date(2018, 11, 24, 20, 34, 58, 651387237, time.UTC),
	//	Data: &models.SensorData{Temperature: 71.0, Humidity: 0.10}})
	//dataset = append(dataset, models.DataPoint{ID: 2, Timestamp: time.Now(), Data: &models.SensorData{Temperature: 74.0, Humidity: 0.23}})

	router := mux.NewRouter()
	router.HandleFunc("/sensordata/api/data", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data/{id}", GetData).Methods("GET")
	router.HandleFunc("/sensordata/api/data", CreateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", UpdateData).Methods("POST")
	router.HandleFunc("/sensordata/api/data/{id}", DeleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

