package handler

import (
	"encoding/json"
	"net/http"

	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

type ProfileResponse struct {
	Id     int64  `json:"id"`
	OpenId string `json:"openid"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := ProfileResponse{
		Id:     user.Id,
		OpenId: user.OpenId,
	}

	logger.Info().Interface("user", user).Msg("user")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
