package model

type Sensor struct {
	IDSensor    int    `gorm:"type:int(11);primary_key;unique_index;autoIncrement:true;column:id_sensor" json:"idSensor"`
	Mac         string `gorm:"type:varchar(17);not null;column:mac" json:"mac"`
	Description string `gorm:"type:varchar(255);column:description" json:"description"`
	Password    string `gorm:"type:varchar(16);not null;column:password" json:"password"`
}

func (Sensor) TableName() string {
	return "sensors"
}
