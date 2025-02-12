package middleware

import (
	"context"
	"net/http"
	"time"

	"wx-miniprogram-backend/internal/log"

	"github.com/rs/zerolog"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

const loggerKey contextKey = "logger"

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.Logger.With().Str("method", r.Method).Str("path", r.URL.Path).Logger()

		start := time.Now()

		// 记录请求开始
		logger.Info().
			Str("remote_addr", r.RemoteAddr).
			Str("user_agent", r.UserAgent()).
			Msg("Request started")

		// 包装 ResponseWriter 以捕获状态码和响应大小
		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK, // 默认状态码
		}

		// 处理请求

		ctx := context.WithValue(r.Context(), loggerKey, &logger)
		next.ServeHTTP(wrapped, r.WithContext(ctx))

		// 计算处理时间
		duration := time.Since(start)

		// 记录请求完成
		logger.Info().
			Int("status", wrapped.status).
			Int("size", wrapped.size).
			Dur("duration", duration).
			Msg("Request completed")
	})
}

func GetLogger(r *http.Request) *zerolog.Logger {
	return r.Context().Value(loggerKey).(*zerolog.Logger)
}
