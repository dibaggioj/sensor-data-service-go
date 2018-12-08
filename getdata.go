package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"errors"
	"fmt"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode(dataset)
	} else {
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		data, err := GetDataPoint(id)
		if err.Error == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		}
	}
}

func GetDataPoint(id int64) (DataPoint, Error) {
	var data DataPoint
	var err Error
	for _, item := range dataset {
		if item.ID == id {
			return item, err
		}
	}
	return data, Error{Error: errors.New("data point not found"),
	Message: fmt.Sprintf("Data point with ID %d not found", id)}
}

func GetDataPointReference(id int64) (*DataPoint, Error) {
	var data *DataPoint
	var err Error
	for index, item := range dataset {
		if item.ID == id {
			return &dataset[index], err
		}
	}
	return data, Error{Error: errors.New("data point not found"),
		Message: fmt.Sprintf("Data point with ID %d not found", id)}
}