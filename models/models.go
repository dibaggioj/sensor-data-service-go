package models

import (
	"time"
)

type DataPoint struct {
	ID        int64       `json:"id,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Data      *SensorData `json:"data,omitempty"`
}

type SensorData struct {
	ID        int64       `json:"id,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Humidity 	float64 `json:"humidity,omitempty"`
}

type DataChangePayload struct {
	ID			int64   	`json:"id,omitempty"`
	Message 	string `json:"message,omitempty"`
}
type Error struct {
	Error 		error `json:"-"`
	Message 	string `json:"message,omitempty"`
}

func (dataPoint *DataPoint) isValid() bool {
	if dataPoint.Data == nil {
		return false
	}
	return true
}