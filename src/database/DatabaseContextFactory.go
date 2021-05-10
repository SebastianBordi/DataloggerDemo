package database

import (
	"log"
)

type DatabaseIdentifier string

const (
	MySQL      DatabaseIdentifier = "mysql"
	PostgreSQL DatabaseIdentifier = "posgres"
	MsSQL      DatabaseIdentifier = "sqlServer"
)

type databaseContextFactory struct {
	DatabaseType DatabaseIdentifier
}

var factoryContext *databaseContextFactory

func GetDataContextFactory() *databaseContextFactory {
	if factoryContext == nil {
		factoryContext = &databaseContextFactory{
			DatabaseType: MySQL,
		}
	}
	return factoryContext
}

func (dcf databaseContextFactory) GetDataContext() IContext {
	if dcf.DatabaseType == PostgreSQL {
		context, err := getPostgreSqlInstance()
		if err != nil {
			log.Panic(err)
		}
		return context
	} else if dcf.DatabaseType == MsSQL {
		return nil
	} else {
		context, err := getMySqlInstance()
		if err != nil {
			log.Panic(err)
		}
		return context
	}
}
