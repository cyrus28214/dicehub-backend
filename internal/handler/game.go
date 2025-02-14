package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

// ListGamesHandler 获取游戏列表
func ListGamesHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	// 获取用户ID
	userId, ok := middleware.GetUserID(r)
	if !ok {
		logger.Warn().Msg("User ID is not set")
		userId = 0
	}

	// 解析标签过滤参数
	tagIds := r.URL.Query().Get("tagIds") // tagIds 以逗号分隔
	var tagIdsList []int64
	if tagIds != "" {
		// 将逗号分隔的字符串转换为整数切片
		tagIdsStrList := strings.Split(tagIds, ",")
		for _, idStr := range tagIdsStrList {
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				logger.Warn().Str("tagId", idStr).Msg("Invalid tag ID format")
				continue
			}
			tagIdsList = append(tagIdsList, id)
		}
	}
	logger.Info().Interface("tagIdsList", tagIdsList).Msg("tagIdsList")

	games, err := model.GetGames(tagIdsList, userId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get games")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

// GetGameHandler 获取游戏详情
func GetGameHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	// 获取用户ID
	userId, ok := middleware.GetUserID(r)
	if !ok {
		logger.Warn().Msg("User ID is not set")
		userId = 0
	}

	// 从URL中获取游戏ID
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	game, err := model.GetGameById(id, userId)
	if err != nil {
		logger.Error().Err(err).Int64("id", id).Msg("Failed to get game")
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
