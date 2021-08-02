package model

type Sensor struct {
	//IDSensor    int    `gorm:"type:int(11);primaryKey;unique;not null;column:id_sensor;autoIncrement" json:"idSensor"`
	ID          uint   `json:"id"`
	Mac         string `gorm:"type:varchar(17);not null;column:mac" json:"mac"`
	Description string `gorm:"type:varchar(255);column:description" json:"description"`
	Password    string `gorm:"type:varchar(16);not null;column:password" json:"password"`
}

func (Sensor) TableName() string {
	return "sensors"
}
