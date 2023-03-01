package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

var config *Config

func Init() {
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$GOPATH/src/github.com/gerardva/go-api/")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&config)
}

func GetConfig() *Config {
	return config
}