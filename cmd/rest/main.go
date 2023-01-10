package main

import (
	rest "github.com/Shteyd/notes-backend/internal/app"
	"github.com/Shteyd/notes-backend/pkg/database"
	"github.com/Shteyd/notes-backend/pkg/env"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	loadLogger()
	loadConfigs()
	loadDatabase()
	rest.Run()
}

func loadLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func loadDatabase() {
	cfg := env.GetEnvConfig()
	if err := database.SetPostgresDatabase(database.PostgresConfig{
		Host:     cfg.DbConfig.Host,
		Port:     cfg.DbConfig.Port,
		Username: cfg.DbConfig.Username,
		Password: cfg.DbConfig.Password,
		DBName:   cfg.DbConfig.DBName,
		SSLMode:  cfg.DbConfig.SSLMode,
	}); err != nil {
		logrus.Fatalf("error occured while initialize database: %s", err.Error())
	}
}

func loadConfigs() {
	if err := env.InitEnv(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	env.SetEnvConfig()
}
