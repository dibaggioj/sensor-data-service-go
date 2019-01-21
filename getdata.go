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
		// TODO: replace with db query
		data, err := GetDataPoint(uint(id))
		if err.Error == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		}
	}
}

func GetDataPoint(id uint) (models.DataPoint, models.Error) {
	// TODO: replace with db query
	var data models.DataPoint
	var err models.Error
	for _, item := range dataset {
		if item.ID == id {
			return item, err
		}
	}
	return data, models.Error{Error: errors.New("data point not found"),
	Message: fmt.Sprintf("Data point with ID %d not found", id)}
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