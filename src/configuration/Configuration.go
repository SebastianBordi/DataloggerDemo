package configuration

import (
	"os"
)

type Config struct {
	database *DatabaseConf
	baseURL  string //base URL for all endpoints
	urlPort  string //exposed port for API

}

type DatabaseConf struct {
	host     string
	port     string
	database string
	user     string
	password string
}

var instance *Config

func GetInstance() (*Config, error) {
	if instance == nil {
		instance = &Config{}
		instance.baseURL = os.Getenv("BASE_URL")
		instance.urlPort = os.Getenv("PORT")

		instance.database = &DatabaseConf{}
		instance.database.host = os.Getenv("DATABASE_HOST")
		instance.database.port = os.Getenv("DATABASE_PORT")
		instance.database.database = os.Getenv("DATABASE_NAME")
		instance.database.user = os.Getenv("DATABASE_USER")
		instance.database.password = os.Getenv("DATABASE_PASSWORD")
	}
	return instance, nil
}

// func GetDatabaseConf() (*DatabaseConf, error) {
// 	var err error
// 	if instance == nil {
// 		_, err = GetInstance()
// 		return nil, err
// 	}
// 	databaseConf := instance.GetDatabaseConf()
// 	return databaseConf, err
// }

//config struct methods
func (c Config) GetBaseURL() string {
	return c.baseURL
}
func (c Config) GetURLPort() string {
	return c.urlPort
}
func (c Config) GetDatabaseConf() *DatabaseConf {
	return c.database
}

//databaseConf struct methods
func (dbc DatabaseConf) GetHost() string {
	return dbc.host
}
func (dbc DatabaseConf) GetPort() string {
	return dbc.port
}
func (dbc DatabaseConf) GetDatabase() string {
	return dbc.database
}
func (dbc DatabaseConf) GetUser() string {
	return dbc.user
}
func (dbc DatabaseConf) GetPassword() string {
	return dbc.password
}
