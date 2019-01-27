package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/dibaggioj/sensor-api/exceptions"
)

type DataPoint struct {
	gorm.Model
	Timestamp  time.Time  	`json:"timestamp"`
	SensorData SensorData	`gorm:"foreignkey:ID" json:"data"`
}
type SensorData struct {
	gorm.Model
	Humidity		float64	`gorm:"type:float(8);" json:"humidity"`
	Temperature		float64 `gorm:"type:float(8);" json:"temperature"`
	TemperatureUnit	rune 	`gorm:"char(1);" json:"temperatureUnit"`
}

type DataChangePayload struct {
	ID			uint   		`json:"id,omitempty"`
	Message 	string 		`json:"message,omitempty"`
}
type Error struct {
	Error 		error 		`json:"error"`
	Message 	string 		`json:"message,omitempty"`
}

func (dataPoint *DataPoint) Validate() error {
	var sensorData *SensorData
	sensorData = &dataPoint.SensorData
	if sensorData == nil {
		return &exceptions.DataValidationError{Reason: "Missing sensor data"}
	} else if sensorData.Humidity < 0 {
		return &exceptions.DataValidationError{Reason: "Humidity cannot be negative"}
	} else if sensorData.TemperatureUnit != 'F' && sensorData.TemperatureUnit != 'C' {
		return &exceptions.DataValidationError{Reason: "Missing temperature units"}
	} else if sensorData.TemperatureUnit == 'F' && sensorData.Temperature < -459.67 {
		return &exceptions.DataValidationError{Reason: "Temperature cannot be below absolute zero"}
	} else if sensorData.TemperatureUnit == 'C' && sensorData.Temperature < -273.15 {
		return &exceptions.DataValidationError{Reason: "Temperature cannot be below absolute zero"}
	}
	return nil
}


