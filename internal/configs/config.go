package configs

type HTTP struct {
	Host     string
	Port     string
	Protocol string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Config struct {
	HTTP                HTTP
	Database            Database
	Environment         string
	DatabaseOnlineStore Database
}

var AppConfig *Config
