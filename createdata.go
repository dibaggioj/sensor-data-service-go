package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/dibaggioj/sensor-api/models"
	"fmt"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	var data models.DataPoint
	_ = json.NewDecoder(r.Body).Decode(&data)

	dataErr := data.Validate()
	if dataErr == nil {
		db.NewRecord(data)	// returns `true` as primary key is blank
		db.Create(&data)
		db.NewRecord(data)	// returns `false` since record has been created

		responsePayload := models.DataChangePayload{ID: data.ID, Message: "Created data point"}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responsePayload)
	} else {
		errPayload := models.Error{Error: dataErr, Message: "Unable to create data point"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
	}
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
		var responsePayload *models.DataChangePayload
		dataErr := data.Validate()
		if dataErr == nil {
			sensorData.Temperature = updatedSensorData.Temperature
			sensorData.TemperatureUnit = updatedSensorData.TemperatureUnit
			sensorData.Humidity = updatedSensorData.Humidity
			db.Save(&data)
			responsePayload = &models.DataChangePayload{ID: data.ID, Message: "Updated data point"}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(responsePayload)
		} else {
			errPayload := models.Error{Error: dataErr, Message: fmt.Sprintf("Unable to update data point with ID %d", id)}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errPayload)
		}
	} else {
		errPayload := models.Error{Error: err, Message: fmt.Sprintf("Unable to find data point with ID %d", id)}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errPayload)
	}
}