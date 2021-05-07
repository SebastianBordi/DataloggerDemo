package main

import (
	"log"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/router"
	"github.com/sebastianbordi/DataloggerDemo/server"
)

func main() {
	conf, err := configuration.GetInstance()
	if err != nil {
		log.Panic(err)
	}
	configureDatabase(conf.GetDatabaseConf())
	rtr := router.GetRouter()
	srv := server.GetServer(conf.GetURLPort())

	srv.Handler = rtr

}

func configureDatabase(config *configuration.DatabaseConf) {
	dbContext, err := database.GetInstance()
	if err != nil {
		log.Panic(err)
	}
	dbContext.SetHost(config.GetHost())
	dbContext.SetPort(config.GetPort())
	dbContext.SetDatabase(config.GetDatabase())
	dbContext.SetUser(config.GetUser())
	dbContext.SetPassword(config.GetPassword())

	err = dbContext.Initialize()
	if err != nil {
		log.Panic(err)
	}
}
