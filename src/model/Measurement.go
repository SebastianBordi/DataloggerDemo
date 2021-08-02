package model

import (
	"time"
)

type Measurement struct {
	ID uint `json:"id"`
	//IDMeasurement int       `gorm:"type:int(11);primaryKey;unique_index;autoIncrement;not null;column:id_measurement" json:"idMeasurement"`
	Temperature float64   `gorm:"type:decimal(5.2);not null;column:temperature" json:"temperature"`
	Humidity    int       `gorm:"type:int(5);not null;column:humidity" json:"humidity"`
	Datetime    time.Time `gorm:"type:datetime;not null;column:date" json:"date"`
	IDSensor    uint      `gorm:"type:int(11);column:id_sensor" json:"-"`
	Sensor      Sensor    `gorm:"foreignKey:IDSensor" json:"sensor"`
}

func (Measurement) TableName() string {
	return "measurements"
}
