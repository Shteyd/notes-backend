package main

import (
	rest "github.com/Shteyd/notes-backend/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	rest.Run()
}
