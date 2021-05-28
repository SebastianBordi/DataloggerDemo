package configuration

import (
	"os"
	"strings"
)

const (
	DEVELOPMENT = "dev"
	PRODUCTION  = "prod"
)

type Config struct {
	database       *DatabaseConf
	baseURL        string //base URL for all endpoints
	urlPort        string //exposed port for API
	tokenKey       string //key fot token decoding
	environment    string
	allowedOrigins []string
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
		instance.tokenKey = os.Getenv("TOKEN_KEY")
		instance.environment = os.Getenv("APP_ENV")
		instance.allowedOrigins = getAllowedOrigins()

		instance.database = &DatabaseConf{}
		instance.database.host = os.Getenv("DATABASE_HOST")
		instance.database.port = os.Getenv("DATABASE_PORT")
		instance.database.database = os.Getenv("DATABASE_NAME")
		instance.database.user = os.Getenv("DATABASE_USER")
		instance.database.password = os.Getenv("DATABASE_PASSWORD")
	}
	return instance, nil
}
func getAllowedOrigins() []string {
	origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ";")
	return origins
}

//config struct methods
func (c Config) GetBaseURL() string {
	return c.baseURL
}
func (c Config) GetURLPort() string {
	return c.urlPort
}
func (c Config) GetTokenKey() string {
	return c.tokenKey
}
func (c Config) GetEnvironment() string {
	var env string
	envLower := strings.ToLower(c.environment)
	if strings.Contains(envLower, "dev") {
		env = DEVELOPMENT
	} else {
		env = PRODUCTION
	}
	return env
}
func (c Config) GetAllowedOrigins() []string {
	return c.allowedOrigins
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
