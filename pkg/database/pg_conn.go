package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func SetPostgresDatabase(cfg PostgresConfig) error {
	var err error

	db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

func GetDatabase() *sqlx.DB {
	return db
}
