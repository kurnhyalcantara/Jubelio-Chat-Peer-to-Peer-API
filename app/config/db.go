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
	DB_SSLMODE string
	DB_MAXOPENCONN int
	DB_MAXIDLECONN int
	DB_MAXLIFETIMECONN int
}

func LoadDBConfig() (*DBConfig, error) {
	dbHost, foundDBHost := os.LookupEnv("DB_HOST")
	dbPort, foundDBPort := os.LookupEnv("DB_PORT")
	dbUsername, foundDBUsername := os.LookupEnv("DB_USERNAME")
	dbPass, foundDBPass := os.LookupEnv("DB_PASS")
	dbName, foundDBName := os.LookupEnv("DB_NAME")
	dbSslMode, foundDBSslMode := os.LookupEnv("DB_SSLMODE")
	dbMaxOpenConn, foundDBMaxOpenConn := os.LookupEnv("DB_MAXOPENCONN")
	dbMaxIdleConn, foundDBMaxIdleConn := os.LookupEnv("DB_MAXIDLECONN")
	dbMaxLifeTimeConn, foundDBMaxLifeTimeConn := os.LookupEnv("DB_MAXLIFETIMECONN")

	if !foundDBHost || !foundDBPort || !foundDBUsername || !foundDBPass || !foundDBName || foundDBSslMode || foundDBMaxOpenConn || foundDBMaxLifeTimeConn || foundDBMaxIdleConn {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.local")
		viper.SetConfigType("env")

		dbHost = viper.GetString("DB_HOST")
		dbPort = viper.GetString("DB_PORT")
		dbUsername = viper.GetString("DB_USERNAME")
		dbPass = viper.GetString("DB_PASS")
		dbName = viper.GetString("DB_NAME")
		dbSslMode = viper.GetString("DB_SSLMODE")
		dbMaxOpenConn = viper.GetString("DB_MAXOPENCONN")
		dbMaxIdleConn = viper.GetString("DB_MAXIDLECONN")
		dbMaxLifeTimeConn = viper.GetString("DB_MAXLIFETIMECONN")
	}

	dbPortInt, _ := strconv.Atoi(dbPort)
	dbMaxOpenConnInt, _ := strconv.Atoi(dbMaxOpenConn)
	dbMaxIdleConnInt, _  := strconv.Atoi(dbMaxIdleConn)
	dbMaxLifeTimeConnInt, _ := strconv.Atoi(dbMaxLifeTimeConn)

	return &DBConfig{
		DB_HOST: dbHost,
		DB_PORT: dbPortInt,
		DB_USERNAME: dbUsername,
		DB_PASS: dbPass,
		DB_NAME: dbName,
		DB_SSLMODE: dbSslMode,
		DB_MAXOPENCONN: dbMaxOpenConnInt,
		DB_MAXIDLECONN: dbMaxIdleConnInt,
		DB_MAXLIFETIMECONN: dbMaxLifeTimeConnInt,
	}, nil
}