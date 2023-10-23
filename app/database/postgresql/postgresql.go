package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kurnhyalcantara/Jubelio-Chat-Peer-to-Peer-API/app/config"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var postgresql = &DB{}

func (db *DB) connect(dbConfig *config.DBConfig) (err error) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_USERNAME,
		dbConfig.DB_PASS,
		dbConfig.DB_NAME,
		dbConfig.DB_SSLMODE,
	)

	postgresql.DB, err = sql.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(dbConfig.DB_MAXOPENCONN)
	db.SetMaxIdleConns(dbConfig.DB_MAXIDLECONN)
	db.SetConnMaxLifetime(time.Duration(dbConfig.DB_MAXLIFETIMECONN))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("can't send ping to database: %v", err)
	}

	return nil
}

func ConnectDB() error {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		return err 
	}
	return postgresql.connect(dbConfig)
}