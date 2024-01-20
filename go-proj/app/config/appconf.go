package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type AppConf struct {
	AppEnvironment string
}

func Init() (*AppConf, error) {
	err := godotenv.Load(".env")

	if err != nil {
		slog.Error("failed when reading .env file")
	}

	config := &AppConf{
		AppEnvironment: os.Getenv("APP_ENV"),
	}

	return config, nil
}
