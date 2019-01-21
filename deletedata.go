package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"errors"
	"fmt"
	"github.com/dibaggioj/sensor-api/models"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	for index, item := range dataset {
		if item.ID == uint(id) {
			// TODO: replace with db query
			dataset = append(dataset[:index], dataset[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Error{Error: errors.New("data point not found"),
		Message: fmt.Sprintf("Unable to delete data point, ID %d not found", id)})
}
