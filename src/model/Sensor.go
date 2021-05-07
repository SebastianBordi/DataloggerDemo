package model

type Sensor struct {
	IDSensor    int
	Mac         string
	Description string
	Password    string
}

func (Sensor) TableName() string {
	return "sensors"
}
