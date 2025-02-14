package handler

import (
	"encoding/json"
	"net/http"

	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

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

	logger.Info().Interface("user", user).Msg("user")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)

	if !ok {
		logger.Error().Msg("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var request struct {
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = model.UpdateUser(userId, request.Avatar, request.Name)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to update user")
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
