package main

import (
	"fmt"
	"net/http"

	"wx-miniprogram-backend/internal/config"
	"wx-miniprogram-backend/internal/handler"
	"wx-miniprogram-backend/internal/log"
)

func main() {
	http.HandleFunc("/api/login", handler.LoginHandler)

	serverAddr := fmt.Sprintf(":%s", config.Cfg.ServerPort)
	log.Logger.Info().Msgf("Server starting on port %s...", config.Cfg.ServerPort)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
