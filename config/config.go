package config

import (
	"github.com/spf13/viper"
)

type App struct {
	Name           string `json:"name"`
	HTTPPort       string `json:"http_port"`
	LogLevel       int    `json:"log_level"`
	LogTimeFormat  string `json:"time_format"`
	ContextTimeout int    `json:"context_timeout "`
}

type MysqlDB struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Driver   string `json:"driver"`
}

type Config struct {
	App     App
	MysqlDB MysqlDB
}

func init() {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func NewConfig() Config {
	return Config{
		App: App{
			Name:           viper.GetString("APP_NAME"),
			HTTPPort:       viper.GetString("PORT"),
			LogLevel:       viper.GetInt("LOG_LEVEL"),
			ContextTimeout: viper.GetInt("CONTEXT_TIMEOUT"),
		},
		MysqlDB: MysqlDB{
			Name:     viper.GetString("MYSQL_DB_NAME"),
			Host:     viper.GetString("MYSQL_DB_HOST"),
			User:     viper.GetString("MYSQL_DB_USER"),
			Port:     viper.GetString("MYSQL_DB_PORT"),
			Password: viper.GetString("MYSQL_DB_PASS"),
			Driver:   viper.GetString("DRIVER"),
		},
	}
}
