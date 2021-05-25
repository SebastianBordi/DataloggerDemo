package model

type User struct {
	ID       uint
	User     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(16)"`
}

func (User) TableName() string {
	return "users"
}
