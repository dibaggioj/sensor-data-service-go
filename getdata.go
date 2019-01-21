package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"errors"
	"fmt"
	"github.com/dibaggioj/sensor-api/models"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		var dataPoints []models.DataPoint
		db.Find(&dataPoints)	// Get all records - SELECT * FROM data_points;
		json.NewEncoder(w).Encode(dataPoints)
	} else {
		id, _ := strconv.ParseUint(params["id"], 10, 64)
		var dataPoint models.DataPoint
		err := db.First(&dataPoint, id).Error
		if err != nil {
			errPayload := models.Error{Error: err, Message: fmt.Sprintf("Unable to find data point with ID %d", id)}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errPayload)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dataPoint)
		}
	}
}

func GetDataPointReference(id uint) (*models.DataPoint, models.Error) {
	var data *models.DataPoint
	var err models.Error
	for index, item := range dataset {
		if item.ID == id {
			return &dataset[index], err
		}
	}
	return data, models.Error{Error: errors.New("data point not found"),
		Message: fmt.Sprintf("Data point with ID %d not found", id)}
}