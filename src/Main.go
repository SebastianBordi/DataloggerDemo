package main

import (
	"log"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/controller"
	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/model"
	"github.com/sebastianbordi/DataloggerDemo/router"
	"github.com/sebastianbordi/DataloggerDemo/server"
)

func main() {
	conf, err := configuration.GetInstance()
	if err != nil {
		log.Panic(err)
	}
	configureDatabase(conf.GetDatabaseConf())
	initControllers()
	rtr := router.GetRouter(conf.GetBaseURL())
	srv := server.GetServer(conf.GetURLPort())

	srv.Handler = rtr
	log.Printf("Listen and serve at :%s", conf.GetURLPort())
	log.Fatal(srv.ListenAndServe())
}

func configureDatabase(config *configuration.DatabaseConf) {
	dataContextFactory := database.GetDataContextFactory()

	dbContext := dataContextFactory.GetDataContext()
	dbContext.SetHost(config.GetHost())
	dbContext.SetPort(config.GetPort())
	dbContext.SetDatabase(config.GetDatabase())
	dbContext.SetUser(config.GetUser())
	dbContext.SetPassword(config.GetPassword())

	err := dbContext.Initialize()
	if err != nil {
		log.Panic(err)
	}

	dbContext.GetContext().AutoMigrate(&model.Sensor{})
	dbContext.GetContext().AutoMigrate(&model.Measurement{})
	dbContext.GetContext().AutoMigrate(&model.User{})
}

func initControllers() {
	dataContextFactory := database.GetDataContextFactory()
	context := dataContextFactory.GetDataContext()

	sensorController := controller.GetSensorController()
	sensorController.InitSensorController(context)
	temperatureController := controller.GetMeasurementController()
	temperatureController.InitMeasurementController(context)
	userController := controller.GetUserController()
	userController.InitUserController(context)
}
