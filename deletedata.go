package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/dibaggioj/sensor-api/models"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	var err error
	var errPayload models.Error
	var dataPoint models.DataPoint
	err = db.First(&dataPoint, id).Error
	if err != nil {
		errPayload = models.Error{Error: err, Message: fmt.Sprintf("Unable to find data point with ID %d", id)}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errPayload)
	} else {
		err = db.Delete(&dataPoint, id).Error
		if err != nil {
			errPayload = models.Error{
				Error:   err,
				Message: fmt.Sprintf("An error occurred while attempting to delete data point with ID %d", id)}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errPayload)
		} else {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(dataPoint)
		}
	}
}
