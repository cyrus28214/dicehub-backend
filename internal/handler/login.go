package handler

import (
	"encoding/json"
	"net/http"

	"wx-miniprogram-backend/internal/log"
	"wx-miniprogram-backend/internal/weixin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error().Err(err).Msg("Failed to decode login request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	result, err := weixin.Code2Session(req.Code)
	if err != nil {
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		OpenId:     result.OpenId,
		SessionKey: result.SessionKey,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
