package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
)

type CreateCommentRequest struct {
	GameId  int64   `json:"game_id"`
	Content string  `json:"content"`
	Rating  float64 `json:"rating"`
}

type UpdateCommentRequest struct {
	Id      int64   `json:"id"`
	Content string  `json:"content"`
	Rating  float64 `json:"rating"`
}

// CreateCommentHandler 创建评论
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error().Err(err).Msg("Failed to decode request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 验证评分范围
	if req.Rating < 0 || req.Rating > 10 {
		http.Error(w, "Rating must be between 0 and 10", http.StatusBadRequest)
		return
	}

	// 检查是否已经评论过
	existingComment, err := model.GetUserGameComment(userId, req.GameId)
	if err == nil && existingComment != nil {
		http.Error(w, "You have already commented on this game", http.StatusBadRequest)
		return
	}

	err = model.CreateComment(userId, req.GameId, req.Content, req.Rating)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetGameCommentsHandler 获取游戏评论列表
func GetGameCommentsHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	gameId, err := strconv.ParseInt(r.URL.Query().Get("game_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	comments, err := model.GetGameComments(gameId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get comments")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// UpdateCommentHandler 更新评论
func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error().Err(err).Msg("Failed to decode request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Rating < 0 || req.Rating > 10 {
		http.Error(w, "Rating must be between 0 and 10", http.StatusBadRequest)
		return
	}

	err := model.UpdateComment(req.Id, userId, req.Content, req.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		} else {
			logger.Error().Err(err).Msg("Failed to update comment")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCommentHandler 删除评论
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	commentId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	err = model.DeleteComment(commentId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		} else {
			logger.Error().Err(err).Msg("Failed to delete comment")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
