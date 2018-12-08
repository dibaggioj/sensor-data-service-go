package main

import "time"

type DataPoint struct {
	ID		int64   	`json:"id,omitempty"`
	Time	time.Time	`json:"time,omitempty"`
	Data	*SensorData `json:"data,omitempty"`
}
type SensorData struct {
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
