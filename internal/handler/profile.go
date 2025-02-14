package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"wx-miniprogram-backend/internal/config"
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

func UpdateUserNameHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var request struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = model.UpdateUserName(userId, request.Name)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to update user name")
		http.Error(w, "Failed to update user name", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLogger(r)

	userId, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 限制文件大小为5MB
	r.ParseMultipartForm(5 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get file")
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 检查文件类型
	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		http.Error(w, "Only JPG/PNG format is supported", http.StatusBadRequest)
		return
	}

	// 创建上传目录
	uploadDir := filepath.Join(config.Cfg.UploadDir, "avatars")
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		logger.Error().Err(err).Msg("Failed to create upload directory")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 生成文件名
	filename := fmt.Sprintf("%d_%d%s",
		userId,
		time.Now().UnixNano(),
		filepath.Ext(header.Filename),
	)

	filepath := filepath.Join(uploadDir, filename)

	// 创建目标文件
	dst, err := os.Create(filepath)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create file")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 保存文件
	if _, err := io.Copy(dst, file); err != nil {
		logger.Error().Err(err).Msg("Failed to save file")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 更新用户头像
	avatarUrl := fmt.Sprintf("%s/uploads/avatars/%s", config.Cfg.BaseUrl, filename) // 暂时这样解决，使用OSS会更好
	if err := model.UpdateUserAvatar(userId, avatarUrl); err != nil {
		logger.Error().Err(err).Msg("Failed to update user avatar")
		http.Error(w, "Failed to update avatar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(avatarUrl))
}
