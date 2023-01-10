package env

import (
	"os"

	"github.com/joho/godotenv"
)

var config *EnvConfig

type EnvConfig struct {
	AppConfig struct {
		Port      string
		Salt      string
		SignInKey string
	}

	DbConfig struct {
		Host     string
		Port     string
		Username string
		DBName   string
		SSLMode  string
		Password string
	}
}

func InitEnv() error {
	return godotenv.Load()
}

func SetEnvConfig() {
	config = &EnvConfig{
		AppConfig: struct {
			Port      string
			Salt      string
			SignInKey string
		}{
			Port:      os.Getenv("APP_HOST"),
			Salt:      os.Getenv("APP_SALT"),
			SignInKey: os.Getenv("APP_SIGNINKEY"),
		},
		DbConfig: struct {
			Host     string
			Port     string
			Username string
			DBName   string
			SSLMode  string
			Password string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("SSL_MODE"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}

func GetEnvConfig() *EnvConfig {
	return config
}
