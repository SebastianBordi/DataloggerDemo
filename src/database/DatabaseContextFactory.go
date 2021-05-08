package database

import (
	"log"
)

type databaseContextFactory struct {
	DatabaseType string
}

var factoryContext *databaseContextFactory

func GetDataContextFactory() *databaseContextFactory {
	if factoryContext == nil {
		factoryContext = &databaseContextFactory{
			DatabaseType: "MySQL",
		}
	}
	return factoryContext
}

func (dcf databaseContextFactory) GetDataContext() IContext {
	if dcf.DatabaseType == "PostgreSQL" {
		return nil
	} else if dcf.DatabaseType == "MsSQL" {
		return nil
	} else {
		context, err := getMySqlInstance()
		if err != nil {
			log.Panic(err)
		}
		return context
	}
}
