package handler

import (
	"encoding/json"
	"net/http"
	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := middleware.GetLogger(r)

	tags, err := model.GetTags()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get tags")
		http.Error(w, "Failed to get tags", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
