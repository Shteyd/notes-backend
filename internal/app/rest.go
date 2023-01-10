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
	cfg := env.GetEnvConfig()
	db := database.GetDatabase()
	repos := repository.NewRepostitory(db)
	services := services.NewService(cfg, repos)
	router := routers.NewHandler(services)

	if err := router.Run(cfg.AppConfig.Port); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
