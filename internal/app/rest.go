package rest

import (
	"github.com/Shteyd/notes-backend/internal/repository"
	"github.com/Shteyd/notes-backend/internal/routers"
	"github.com/Shteyd/notes-backend/internal/services"
	"github.com/Shteyd/notes-backend/pkg/database"
	"github.com/Shteyd/notes-backend/pkg/env"
	"github.com/sirupsen/logrus"
)

func Run() {
	if err := env.InitEnv(); err != nil {
		logrus.Fatalf("error occured while initialize env-config: %s", err.Error())
	}

	cfg := env.GetEnvConfig()

	db, err := database.NewPostgresDatabase(database.PostgresConfig{
		Host:     cfg.DbConfig.Host,
		Port:     cfg.DbConfig.Port,
		Username: cfg.DbConfig.Username,
		Password: cfg.DbConfig.Password,
		DBName:   cfg.DbConfig.DBName,
		SSLMode:  cfg.DbConfig.SSLMode,
	})

	if err != nil {
		logrus.Fatalf("error occured while initialize database: %s", err.Error())
	}
	repos := repository.NewRepostitory(db)
	services := services.NewService(cfg, repos)
	router := routers.NewHandler(services)
	go func() {
		if err := router.Run(cfg.AppConfig.Port); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
}
