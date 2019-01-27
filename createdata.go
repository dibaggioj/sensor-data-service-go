package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/dibaggioj/sensor-api/models"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	var data models.DataPoint
	_ = json.NewDecoder(r.Body).Decode(&data)

	db.NewRecord(data)	// returns `true` as primary key is blank
	db.Create(&data)
	db.NewRecord(data)	// returns `false` since record has been created

	responsePayload := models.DataChangePayload{ID: data.ID, Message: "Created data point"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responsePayload)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	var data models.DataPoint
	err := db.First(&data, id).Error
	if err == nil {
		var updatedData models.DataPoint
		_ = json.NewDecoder(r.Body).Decode(&updatedData)
		updatedData.ID = uint(id)
		var sensorData *models.SensorData
		var updatedSensorData *models.SensorData
		sensorData = &data.SensorData
		updatedSensorData = &updatedData.SensorData
		if updatedSensorData != nil {
			sensorData.Temperature = updatedSensorData.Temperature
			sensorData.Humidity = updatedSensorData.Humidity
		}
		db.Save(&data)
		responsePayload := models.DataChangePayload{ID: data.ID, Message: "Updated data point"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responsePayload)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}