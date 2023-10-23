package config

func LoadConfigs() (*AppConfig, *DBConfig, error) {
	appConfig, errApp := LoadAppConfig()
	if errApp != nil {
		return nil, nil, errApp
	}

	dbConfig, errDb := LoadDBConfig()
	if errDb != nil {
		return nil, nil, errDb
	}

	return appConfig, dbConfig, nil
}