package main

import (
	_ "github.com/lib/pq"	// package needed (for registering its drivers with the database/sql package) but never directly referenced in code
	"github.com/dibaggioj/sensor-api/models"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tkanos/gonfig"
	"fmt"
	"github.com/jinzhu/gorm"
)

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


	db.AutoMigrate(&models.DataPoint{}, &models.SensorData{})

	db = db.Set("gorm:auto_preload", true)

	route()
}

