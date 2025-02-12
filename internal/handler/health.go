package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"wx-miniprogram-backend/internal/database"
	"wx-miniprogram-backend/internal/log"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Database  string `json:"database"`
	Timestamp int64  `json:"timestamp"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status:    "ok",
		Database:  "ok",
		Timestamp: time.Now().Unix(),
	}

	// 检查数据库连接
	err := database.DB.Ping()
	if err != nil {
		log.Logger.Error().Err(err).Msg("Database health check failed")
		response.Database = "error"
		response.Status = "error"
	}

	w.Header().Set("Content-Type", "application/json")
	if response.Status != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	json.NewEncoder(w).Encode(response)
}
