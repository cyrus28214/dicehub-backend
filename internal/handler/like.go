package handler

import (
	"encoding/json"
	"net/http"

	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

// LikeGameHandler 处理游戏点赞
func LikeGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := middleware.GetLogger(r)

	// 获取用户ID
	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 从请求体获取游戏ID
	var requestBody struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	gameId := requestBody.Id
	if gameId == 0 {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	// 检查是否已点赞
	liked, err := model.IsGameLiked(userId, gameId)
	if err != nil {
		logger.Error().Err(err).Int64("game_id", gameId).Msg("Failed to check if game is liked")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if liked {
		http.Error(w, "Game already liked", http.StatusBadRequest)
		return
	}

	// 点赞
	err = model.LikeGame(userId, gameId)
	if err != nil {
		logger.Error().Err(err).Int64("game_id", gameId).Msg("Failed to like game")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 返回ok
	w.Write([]byte("ok"))
}

// UnlikeGameHandler 处理取消点赞
func UnlikeGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := middleware.GetLogger(r)

	// 获取用户ID
	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 从请求体获取游戏ID
	var requestBody struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	gameId := requestBody.Id
	if gameId == 0 {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	// 取消点赞
	if err := model.UnlikeGame(userId, gameId); err != nil {
		logger.Error().Err(err).Int64("game_id", gameId).Msg("Failed to unlike game")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 返回ok
	w.Write([]byte("ok"))
}
