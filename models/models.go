package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"encoding/json"
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

// Set temperature unit rune value from either a string (its first character) or an int
func (sd *SensorData) UnmarshalJSON(data []byte) error {
	type Alias SensorData
	// Auxiliary struct with temperatureUnit as a string, e.g, "F" or "C" or even "Celsius"
	aux := &struct {
		TemperatureUnit string `json:"temperatureUnit"`
		*Alias
	}{
		Alias: (*Alias)(sd),
	}
	type Alias2 SensorData
	aux2 := &struct {
		TemperatureUnit rune `json:"temperatureUnit"`
		*Alias2
	}{
		Alias2: (*Alias2)(sd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		// If there was an error then
		if err := json.Unmarshal(data, &aux2); err != nil {
			return err
		}
		sd.TemperatureUnit = aux2.TemperatureUnit
		return nil
	}
	chars := []rune(aux.TemperatureUnit)
	if len(chars) > 0 {
		sd.TemperatureUnit = chars[0]
	} else {
		sd.TemperatureUnit = 0
	}
	return nil
}

func (dataPoint *DataPoint) Validate() error {
	var sensorData *SensorData
	sensorData = &dataPoint.SensorData
	if sensorData == nil {
		return &exceptions.DataValidationError{Reason: "Missing sensor data"}
	} else if sensorData.Humidity < 0 {
		return &exceptions.DataValidationError{Reason: "Humidity cannot be negative"}
	} else if sensorData.TemperatureUnit == 0 {
		return &exceptions.DataValidationError{Reason: "Missing temperature unit"}
	} else if sensorData.TemperatureUnit != 'F' && sensorData.TemperatureUnit != 'C'&&
		sensorData.TemperatureUnit != 'K' && sensorData.TemperatureUnit != 'R' {
		return &exceptions.DataValidationError{Reason: "Invalid temperature unit. Must be 'C', 'F', 'K', or 'R'"}
	} else if sensorData.TemperatureUnit == 'F' && sensorData.Temperature < -459.67 {
		return &exceptions.DataValidationError{Reason: "Temperature cannot be below absolute zero"}
	} else if sensorData.TemperatureUnit == 'C' && sensorData.Temperature < -273.15 {
		return &exceptions.DataValidationError{Reason: "Temperature cannot be below absolute zero"}
	} else if (sensorData.TemperatureUnit == 'K' || sensorData.TemperatureUnit == 'R') && sensorData.Temperature < 0 {
		return &exceptions.DataValidationError{Reason: "Absolute temperature cannot be negative"}
	}
	return nil
}


