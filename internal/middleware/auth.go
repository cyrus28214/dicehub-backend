package middleware

import (
	"context"
	"net/http"
	"strings"

	"wx-miniprogram-backend/internal/crypto"
	"wx-miniprogram-backend/internal/log"
)

const UserIdKey contextKey = "user_id"

// Auth 验证JWT token并将用户ID注入到context中
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从Authorization header中获取token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Logger.Warn().Msg("No Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 检查Authorization header格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Logger.Warn().Msg("Invalid Authorization header format")
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// 验证token
		// TODO: 处理过期
		claims, err := crypto.VerifyToken(parts[1])
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		log.Logger.Info().Interface("claims", claims).Msg("claims")

		// 将用户ID注入到context中
		ctx := context.WithValue(r.Context(), UserIdKey, claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID 从context中获取用户ID
func GetUserID(r *http.Request) (int64, bool) {
	userId, ok := r.Context().Value(UserIdKey).(int64)
	return userId, ok
}
