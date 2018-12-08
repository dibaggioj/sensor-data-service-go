package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"errors"
	"fmt"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for index, item := range dataset {
		if item.ID == id {
			dataset = append(dataset[:index], dataset[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{Error: errors.New("data point not found"),
		Message: fmt.Sprintf("Unable to delete data point, ID %d not found", id)})
}
