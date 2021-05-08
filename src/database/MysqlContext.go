package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlContext struct {
	context     *gorm.DB
	host        string `default:"127.0.0.1"`
	database    string `default:"datalogger"`
	port        string `default:"3306"`
	user        string `default:"datalogger"`
	password    string `default:"123456"`
	initialized bool
}

var instance *mysqlContext

<<<<<<< HEAD:src/database/MysqlContext.go
func getMySqlInstance() (*mysqlContext, error) {
=======
func GetInstance() (*mysqlContext, error) {
>>>>>>> 4b4af171392f56cdea87281e99db27059bc23977:src/database/MysqlContext.go
	var err error
	if instance == nil {
		instance = &mysqlContext{}
	}

	return instance, err
}

//Initialize with parameters of current databaseContext structure
//Return error from gorm.Open()
func (ctx *mysqlContext) Initialize() error {

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

func (ctx *mysqlContext) GetContext() *gorm.DB {
	return ctx.context
}

func (ctx *mysqlContext) GetConnectionString() string {
	return ctx.user + ":" + ctx.password + "@(" + ctx.host + ":" + ctx.port + ")/" + ctx.database + "?charset=utf8&parseTime=True&loc=Local"
}

func (ctx *mysqlContext) SetHost(host string) {
	ctx.host = host
}
func (ctx *mysqlContext) SetPort(port string) {
	ctx.port = port
}
func (ctx *mysqlContext) SetDatabase(database string) {
	ctx.database = database
}
func (ctx *mysqlContext) SetUser(user string) {
	ctx.user = user
}
func (ctx *mysqlContext) SetPassword(password string) {
	ctx.password = password
}
