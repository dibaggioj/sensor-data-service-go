package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type DataPoint struct {
	gorm.Model
	Timestamp  time.Time   `json:"timestamp"`
	SensorData SensorData `gorm:"foreignkey:ID" json:"data"`
}
type SensorData struct {
	gorm.Model
	Temperature	float64 	`gorm:"type:float(8);ps"`
	Humidity	float64 	`gorm:"type:float(8);"`
}

type DataChangePayload struct {
	ID			uint   	`json:"id,omitempty"`
	Message 	string `json:"message,omitempty"`
}
type Error struct {
	Error 		error `json:"-"`
	Message 	string `json:"message,omitempty"`
}

func (dataPoint DataPoint) IsValid() bool {
	var sensorData *SensorData
	sensorData = &dataPoint.SensorData
	if sensorData == nil {
		return false
	} else if sensorData.Humidity < 0 {
		return false
	} else if sensorData.Temperature < -459.67 {
		return false
	}
	return true
}


