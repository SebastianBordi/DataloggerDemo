package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgrsqlContext struct {
	context     *gorm.DB
	host        string `default:"127.0.0.1"`
	database    string `default:"datalogger"`
	port        string `default:"5432"`
	user        string `default:"datalogger"`
	password    string `default:"123456"`
	initialized bool
}

var psqlInstance *postgrsqlContext

func getPostgreSqlInstance() (*postgrsqlContext, error) {
	var err error
	if psqlInstance == nil {
		psqlInstance = &postgrsqlContext{}
	}

	return psqlInstance, err
}

//Initialize with parameters of current databaseContext structure
//Return error from gorm.Open()
func (ctx *postgrsqlContext) Initialize() error {

	newLogger := logger.New(
		log.New(os.Stdout, "\n\r", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(postgres.Open(ctx.GetConnectionString()), &gorm.Config{Logger: newLogger})
	if err == nil {
		ctx.initialized = true
		ctx.context = db
	} else {
		ctx.initialized = false
		ctx.context = nil
	}
	return err
}

func (ctx *postgrsqlContext) GetContext() *gorm.DB {
	return ctx.context
}

func (ctx *postgrsqlContext) GetConnectionString() string {
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		ctx.host, ctx.user, ctx.password, ctx.database, ctx.port)
	conString += " sslmode=disable TimeZone=UTC"
	return conString
}

func (ctx *postgrsqlContext) SetHost(host string) {
	ctx.host = host
}
func (ctx *postgrsqlContext) SetPort(port string) {
	ctx.port = port
}
func (ctx *postgrsqlContext) SetDatabase(database string) {
	ctx.database = database
}
func (ctx *postgrsqlContext) SetUser(user string) {
	ctx.user = user
}
func (ctx *postgrsqlContext) SetPassword(password string) {
	ctx.password = password
}
