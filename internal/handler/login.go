package handler

import (
	"encoding/json"
	"net/http"

	"wx-miniprogram-backend/internal/crypto"
	"wx-miniprogram-backend/internal/middleware"
	"wx-miniprogram-backend/internal/model"
	"wx-miniprogram-backend/internal/weixin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	OpenId string `json:"openId"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	logger := middleware.GetLogger(r)

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error().Err(err).Msg("Failed to decode login request")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Code == "" {
		logger.Warn().Msg("Code is required")
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	// 调用微信接口获取openid和session_key
	wxResp, err := weixin.Code2Session(req.Code)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get session from WeChat")
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	// 查找或创建用户
	user, err := model.FindOrCreateByOpenId(wxResp.OpenId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to find or create user")
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	// 生成JWT token
	tokenString, err := crypto.SignToken(user.Id)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to generate token")
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Token:  tokenString,
		OpenId: user.OpenId,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
