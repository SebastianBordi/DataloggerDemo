package database

import "gorm.io/gorm"

type IContext interface {
	GetContext() *gorm.DB
}
