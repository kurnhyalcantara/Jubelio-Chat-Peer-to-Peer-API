package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DB_HOST	string
	DB_PORT	int
	DB_USERNAME	string
	DB_PASS	string
	DB_NAME string
}

func LoadDBConfig() (*DBConfig, error) {
	dbHost, foundDBHost := os.LookupEnv("DB_HOST")
	dbPort, foundDBPort := os.LookupEnv("DB_PORT")
	dbUsername, foundDBUsername := os.LookupEnv("DB_USERNAME")
	dbPass, foundDBPass := os.LookupEnv("DB_PASS")
	dbName, foundDBName := os.LookupEnv("DB_NAME")

	if !foundDBHost || !foundDBPort || !foundDBUsername || !foundDBPass || !foundDBName {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.local")
		viper.SetConfigType("env")

		dbHost = viper.GetString("DB_HOST")
		dbPort = viper.GetString("DB_PORT")
		dbUsername = viper.GetString("DB_USERNAME")
		dbPass = viper.GetString("DB_PASS")
		dbName = viper.GetString("DB_NAME")
	}

	dbPortInt, _ := strconv.Atoi(dbPort)

	return &DBConfig{
		DB_HOST: dbHost,
		DB_PORT: dbPortInt,
		DB_USERNAME: dbUsername,
		DB_PASS: dbPass,
		DB_NAME: dbName,
	}, nil
}