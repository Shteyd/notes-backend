package env

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var config *EnvConfig
var once sync.Once

type EnvConfig struct {
	AppConfig struct {
		Port      string `env:"APP_PORT" env-default:"8080"`
		Salt      string `env:"APP_SALT" env-required:"true"`
		SignInKey string `env:"APP_SIGNINKEY" env-required:"true"`
	}

	DbConfig struct {
		Host     string `env:"PG_HOST" env-required:"true"`
		Port     string `env:"PG_PORT" env-required:"true"`
		Username string `env:"PG_USERNAME" env-required:"true"`
		DBName   string `env:"DB_NAME" env-required:"true"`
		SSLMode  string `env:"SSL_MODE" env-default:"false"`
		Password string `env:"PG_PASSWORD" env-required:"true"`
	}
}

func InitEnv() error {
	return godotenv.Load()
}

func LoadEnvConfig() error {
	var err error
	once.Do(func() {
		logrus.Print("gather config")

		if err = cleanenv.ReadEnv(config); err != nil {
			return
		}
	})

	return err
}

func GetEnvConfig() *EnvConfig {
	return config
}
