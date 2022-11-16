package config

import (
	"github.com/rotisserie/eris"
	"github.com/spf13/viper"
)

func Setup() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(eris.Wrap(err, "error loading config"))
		}
	}
	viper.AutomaticEnv()
}

const (
	DatabaseUser     = "DB_USER"
	DatabasePassword = "DB_PASSWORD"
	DatabaseHost     = "DB_HOST"
	DatabasePort     = "DB_PORT"
	DatabaseName     = "DB_NAME"

	RedisDatabase = "REDIS_DATABASE"
	RedisPassword = "REDIS_PASSWORD"
	RedisAddr     = "REDIS_ADDR"

	GoogleApplicationCredentials = "GOOGLE_APPLICATION_CREDENTIALS"
)
