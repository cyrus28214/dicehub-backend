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

	mux := http.NewServeMux()

	mux.Handle("/api/health", middleware.Use(
		http.HandlerFunc(handler.HealthHandler),
		middleware.Logger,
	))

	mux.Handle("/api/profile", middleware.Use(
		http.HandlerFunc(handler.ProfileHandler),
		middleware.Logger,
		middleware.Auth,
	))

	mux.Handle("/api/game/like", middleware.Use(
		http.HandlerFunc(handler.LikeGameHandler),
		middleware.Logger,
		middleware.Auth,
	))

	mux.Handle("/api/game/unlike", middleware.Use(
		http.HandlerFunc(handler.UnlikeGameHandler),
		middleware.Logger,
		middleware.Auth,
	))

	mux.Handle("/api/login", middleware.Use(
		http.HandlerFunc(handler.LoginHandler),
		middleware.Logger,
	))

	mux.Handle("/api/games", middleware.Use(
		http.HandlerFunc(handler.ListGamesHandler),
		middleware.Logger,
		middleware.Auth,
	))

	mux.Handle("/api/game", middleware.Use(
		http.HandlerFunc(handler.GetGameHandler),
		middleware.Logger,
		middleware.Auth,
	))

	mux.Handle("/api/tags", middleware.Use(
		http.HandlerFunc(handler.GetTagsHandler),
		middleware.Logger,
	))

	serverAddr := fmt.Sprintf(":%s", config.Cfg.ServerPort)
	log.Logger.Info().Msgf("Server starting on port %s...", config.Cfg.ServerPort)

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
