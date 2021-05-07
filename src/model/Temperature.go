package model

import "time"

type Temperature struct {
	IDTemperature int
	Temperature   float64
	Datetime      time.Time
	Sensor        Sensor
	IDSensor      int
}

func (Temperature) TableName() string {
	return "Temperatures"
}
