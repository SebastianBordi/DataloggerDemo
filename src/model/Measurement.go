package model

import "time"

type Measurement struct {
	IDMeasurement int       `gorm:"type:int(11);primaryKey;unique;autoIncrement;not null;column:id_measurement" json:"idMeasurement"`
	Temperature   float64   `gorm:"type:decimal(5.2);not null;column:temperature" json:"temperature"`
	Humedity      int       `gorm:"type:int(5);not null;column:humedity" json:"humedity"`
	Datetime      time.Time `gorm:"type:date;not null;column:date" json:"date"`
	Sensor        Sensor    `json:"sensor"`
	IDSensor      int       `gorm:"type:int(11);not null;column:id_sensor" json:"-"`
}

func (Measurement) TableName() string {
	return "Measurements"
}
