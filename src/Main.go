package main

import (
	"log"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/controller"
	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/middleware"
	"github.com/sebastianbordi/DataloggerDemo/model"
	"github.com/sebastianbordi/DataloggerDemo/router"
	"github.com/sebastianbordi/DataloggerDemo/server"
	"github.com/sebastianbordi/DataloggerDemo/socket"
	"github.com/sebastianbordi/DataloggerDemo/view"
)

func main() {
	conf, err := configuration.GetInstance()
	if err != nil {
		log.Panic(err)
	}
	configureDatabase(conf)
	initDependencies()
	rtr := router.GetRouter(conf.GetBaseURL())
	srv := server.GetServer(conf.GetURLPort())

	srv.Handler = middleware.Middleware(rtr)
	//srv.Handler = rtr

	log.Println("Datalogger Demo Backend v0.1-beta")
	log.Printf("Running in %s mode", conf.GetEnvironment())
	log.Printf("Listen and serve at :%s, base api url = %s", conf.GetURLPort(), conf.GetBaseURL())
	log.Fatal(srv.ListenAndServe())
}

func configureDatabase(config *configuration.Config) {
	dataContextFactory := database.GetDataContextFactory()

	dataContextFactory.DatabaseType = database.MySQL

	dbContext := dataContextFactory.GetDataContext()
	dbContext.SetHost(config.GetDatabaseConf().GetHost())
	dbContext.SetPort(config.GetDatabaseConf().GetPort())
	dbContext.SetDatabase(config.GetDatabaseConf().GetDatabase())
	dbContext.SetUser(config.GetDatabaseConf().GetUser())
	dbContext.SetPassword(config.GetDatabaseConf().GetPassword())

	isDevMode := config.GetEnvironment() == configuration.DEVELOPMENT
	err := dbContext.Initialize(isDevMode)
	if err != nil {
		log.Panic(err)
	}

	dbContext.GetContext().AutoMigrate(&model.Sensor{})
	dbContext.GetContext().AutoMigrate(&model.Measurement{})
	dbContext.GetContext().AutoMigrate(&model.User{})
}

func initDependencies() {
	dataContextFactory := database.GetDataContextFactory()
	context := dataContextFactory.GetDataContext()

	sensorController := controller.GetSensorController()
	sensorController.InitSensorController(context)
	temperatureController := controller.GetMeasurementController()
	temperatureController.InitMeasurementController(context)
	userController := controller.GetUserController()
	userController.InitUserController(context)

	socket.BasicResponse = view.BasicResponse
}
