package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"wx-miniprogram-backend/internal/database"
	"wx-miniprogram-backend/internal/middleware"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Database  string `json:"database"`
	Timestamp int64  `json:"timestamp"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ok",
		Database:  "ok",
		Timestamp: time.Now().Unix(),
	}

	logger := middleware.GetLogger(r)

	// 检查数据库连接
	err := database.DB.Ping()
	if err != nil {
		logger.Error().Err(err).Msg("Database health check failed")
		response.Database = "error"
		response.Status = "error"
	}

	w.Header().Set("Content-Type", "application/json")
	if response.Status != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	json.NewEncoder(w).Encode(response)
}
