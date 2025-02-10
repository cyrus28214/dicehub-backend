package config

import (
	"os"

	"wx-miniprogram-backend/internal/log"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	DB         DBConfig
	ServerPort string
}

var Cfg Config

func init() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	SERVER_PORT := os.Getenv("SERVER_PORT")

	if DB_HOST == "" {
		log.Logger.Error().Msg("DB_HOST is required")
	}
	if DB_PORT == "" {
		log.Logger.Error().Msg("DB_PORT is required")
	}
	if DB_USER == "" {
		log.Logger.Error().Msg("DB_USER is required")
	}
	if DB_PASSWORD == "" {
		log.Logger.Error().Msg("DB_PASSWORD is required")
	}
	if DB_NAME == "" {
		log.Logger.Error().Msg("DB_NAME is required")
	}
	if SERVER_PORT == "" {
		SERVER_PORT = "8080"
	}

	Cfg = Config{
		DB: DBConfig{
			Host:     DB_HOST,
			Port:     DB_PORT,
			User:     DB_USER,
			Password: DB_PASSWORD,
			Name:     DB_NAME,
		},
		ServerPort: SERVER_PORT,
	}
}
