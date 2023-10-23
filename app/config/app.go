package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HOST	string
	PORT	int
}

func LoadAppConfig() (*AppConfig, error) {
	host, foundHost := os.LookupEnv("APP_HOST")
	port, foundPort := os.LookupEnv("APP_PORT")

	if !foundHost || !foundPort {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.local")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		host = viper.GetString("APP_HOST")
		port = viper.GetString("APP_PORT")
	}

	appPort, _ := strconv.Atoi(port)
	
	return &AppConfig{
		HOST: host,
		PORT: appPort,
	}, nil
}
