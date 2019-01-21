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

	sensorData := data.Data

	db.NewRecord(sensorData) // => returns `true` as primary key is blank
	db.Create(&sensorData)
	db.NewRecord(sensorData)

	// TODO: hook up foreign key
	db.NewRecord(data) // => returns `true` as primary key is blank
	db.Create(&data)
	db.NewRecord(data)

	responsePayload := models.DataChangePayload{ID: data.ID, Message: "Created data point"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responsePayload)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var data *models.DataPoint
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	// TODO: replace with db query
	data, err := GetDataPointReference(uint(id))
	if err.Error == nil {
		var updatedData models.DataPoint
		_ = json.NewDecoder(r.Body).Decode(&updatedData)
		updatedData.ID = uint(id)
		*data = updatedData
		responsePayload := models.DataChangePayload{ID: data.ID, Message: "Updated data point"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responsePayload)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}

// TODO: replace with DB query
func getNextID() uint {
	return dataset[len(dataset)-1].ID + 1
}