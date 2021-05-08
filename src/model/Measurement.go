package model

import "time"

type Measurement struct {
	IDMeasurement int
	Temperature   float64
	Humedity      int
	Datetime      time.Time
	Sensor        Sensor
	IDSensor      int
}

func (Measurement) TableName() string {
	return "Measurements"
}
