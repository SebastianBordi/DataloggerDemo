package database

import "gorm.io/gorm"

type IContext interface {
	GetContext() *gorm.DB
<<<<<<< HEAD
	SetHost(string)
	SetPort(string)
	SetDatabase(string)
	SetUser(string)
	SetPassword(string)
	Initialize() error
=======
>>>>>>> 4b4af171392f56cdea87281e99db27059bc23977
}
