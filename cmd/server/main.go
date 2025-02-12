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
	database.Connect()
	defer database.Close()

	// 创建一个新的路由器
	mux := http.NewServeMux()

	// 注册路由
	mux.HandleFunc("/api/login", handler.LoginHandler)
	mux.HandleFunc("/api/health", handler.HealthHandler)

	// 创建带中间件的处理器
	handler := middleware.Logger(mux)

	serverAddr := fmt.Sprintf(":%s", config.Cfg.ServerPort)
	log.Logger.Info().Msgf("Server starting on port %s...", config.Cfg.ServerPort)

	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		log.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
