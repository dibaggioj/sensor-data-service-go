package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type DataPoint struct {
	gorm.Model
	//ID			uint     	`gorm:"primary_key" json:"id"`
	Timestamp 	time.Time	`json:"timestamp"`
	Data		*SensorData	`json:"data"`
}
type SensorData struct {
	gorm.Model
	//ID        	uint      `gorm:"primary_key" json:"id"`
	Temperature	float64 	`gorm:"type:float(8);ps" json:"temperature,omitempty"`
	Humidity	float64 	`gorm:"type:float(8);" json:"humidity,omitempty"`
}

//`gorm:"type:varchar(100);unique_index"`

//type DataPoint struct {
//	ID        int64       `json:"id,omitempty"`
//	Timestamp time.Time   `json:"timestamp,omitempty"`
//	Data      *SensorData `json:"data,omitempty"`
//}

//type SensorData struct {
//	ID        int64       `json:"id,omitempty"`
//	Temperature float64 `json:"temperature,omitempty"`
//	Humidity 	float64 `json:"humidity,omitempty"`
//}

type DataChangePayload struct {
	ID			uint   	`json:"id,omitempty"`
	Message 	string `json:"message,omitempty"`
}
type Error struct {
	Error 		error `json:"-"`
	Message 	string `json:"message,omitempty"`
}

func (dataPoint *DataPoint) IsValid() bool {
	if dataPoint.Data == nil {
		return false
	}
	return true
}


