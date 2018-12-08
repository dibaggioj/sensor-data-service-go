package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	var data DataPoint
	_ = json.NewDecoder(r.Body).Decode(&data)
	data.ID = getNextID()

	dataset = append(dataset, data)
	responsePayload := DataChangePayload{ID: data.ID, Message: "Created data point"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responsePayload)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var data *DataPoint
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	data, err := GetDataPointReference(id)
	if err.Error == nil {
		var updatedData DataPoint
		_ = json.NewDecoder(r.Body).Decode(&updatedData)
		updatedData.ID = id
		*data = updatedData
		responsePayload := DataChangePayload{ID: data.ID, Message: "Updated data point"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responsePayload)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}

// TODO: replace with DB query
func getNextID() int64 {
	return dataset[len(dataset)-1].ID + 1
}