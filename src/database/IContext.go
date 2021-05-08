package database

import "gorm.io/gorm"

type IContext interface {
	GetContext() *gorm.DB
	SetHost(string)
	SetPort(string)
	SetDatabase(string)
	SetUser(string)
	SetPassword(string)
	Initialize() error
}
