package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
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