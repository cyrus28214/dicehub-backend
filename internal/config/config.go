package config

import (
	"os"

	"wx-miniprogram-backend/internal/log"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type WeixinConfig struct {
	AppId     string
	AppSecret string
}

type Config struct {
	Database   DatabaseConfig
	ServerPort string
	Weixin     WeixinConfig
	JwtSecret  []byte
}

var Cfg Config

func init() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	WX_APP_ID := os.Getenv("WX_APP_ID")
	WX_APP_SECRET := os.Getenv("WX_APP_SECRET")
	JWT_SECRET := os.Getenv("JWT_SECRET")

	if DB_HOST == "" {
		log.Logger.Fatal().Msg("DB_HOST is required")
	}
	if DB_PORT == "" {
		log.Logger.Fatal().Msg("DB_PORT is required")
	}
	if DB_USER == "" {
		log.Logger.Fatal().Msg("DB_USER is required")
	}
	if DB_PASSWORD == "" {
		log.Logger.Fatal().Msg("DB_PASSWORD is required")
	}
	if DB_NAME == "" {
		log.Logger.Fatal().Msg("DB_NAME is required")
	}
	if SERVER_PORT == "" {
		SERVER_PORT = "8080"
	}
	if WX_APP_ID == "" {
		log.Logger.Fatal().Msg("WX_APP_ID is required")
	}
	if WX_APP_SECRET == "" {
		log.Logger.Fatal().Msg("WX_APP_SECRET is required")
	}
	if JWT_SECRET == "" {
		log.Logger.Fatal().Msg("JWT_SECRET is required")
	}

	Cfg = Config{
		Database: DatabaseConfig{
			Host:     DB_HOST,
			Port:     DB_PORT,
			User:     DB_USER,
			Password: DB_PASSWORD,
			Name:     DB_NAME,
		},
		ServerPort: SERVER_PORT,
		Weixin: WeixinConfig{
			AppId:     WX_APP_ID,
			AppSecret: WX_APP_SECRET,
		},
		JwtSecret: []byte(JWT_SECRET),
	}
}
