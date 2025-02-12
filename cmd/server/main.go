package main

import (
	"fmt"
	"net/http"

	"wx-miniprogram-backend/internal/config"
	"wx-miniprogram-backend/internal/database"
	"wx-miniprogram-backend/internal/handler"
	"wx-miniprogram-backend/internal/log"
	"wx-miniprogram-backend/internal/middleware"
)

func main() {
	// 连接数据库

	log.Logger.Info().Msg("Connecting to database...")

	database.Connect()
	defer database.Close()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/api/profile", handler.ProfileHandler)
	authHandler := middleware.Auth(authMux)

	logMux := http.NewServeMux()
	logMux.HandleFunc("/api/login", handler.LoginHandler)
	logMux.HandleFunc("/api/health", handler.HealthHandler)
	logMux.HandleFunc("/api/games", handler.ListGamesHandler)
	logMux.HandleFunc("/api/game", handler.GetGameHandler)
	logMux.Handle("/api/profile", authHandler)
	logHandler := middleware.Logger(logMux)

	handler := logHandler

	serverAddr := fmt.Sprintf(":%s", config.Cfg.ServerPort)
	log.Logger.Info().Msgf("Server starting on port %s...", config.Cfg.ServerPort)

	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		log.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
