package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseContext struct {
	context     *gorm.DB
	host        string `default:"127.0.0.1"`
	database    string `default:"datalogger"`
	port        string `default:"3306"`
	user        string `default:"datalogger"`
	password    string `default:"123456"`
	initialized bool
}

var instance *databaseContext

func GetInstance() (*databaseContext, error) {
	var err error
	if instance == nil {
		instance = &databaseContext{}
	}

	return instance, err
}

//Initialize with parameters of current databaseContext structure
//Return error from gorm.Open()
func (ctx *databaseContext) Initialize() error {

	db, err := gorm.Open(mysql.Open(ctx.GetConnectionString()), &gorm.Config{})
	if err == nil {
		ctx.initialized = true
		ctx.context = db
	} else {
		ctx.initialized = false
		ctx.context = nil
	}
	return err
}

func (ctx *databaseContext) GetContext() *gorm.DB {
	return ctx.context
}

func (ctx *databaseContext) GetConnectionString() string {
	return ctx.user + ":" + ctx.password + "@(" + ctx.host + ":" + ctx.port + ")/" + ctx.database + "?charset=utf8&parseTime=True&loc=Local"
}

func (ctx *databaseContext) SetHost(host string) {
	ctx.host = host
}
func (ctx *databaseContext) SetPort(port string) {
	ctx.port = port
}
func (ctx *databaseContext) SetDatabase(database string) {
	ctx.database = database
}
func (ctx *databaseContext) SetUser(user string) {
	ctx.user = user
}
func (ctx *databaseContext) SetPassword(password string) {
	ctx.password = password
}
