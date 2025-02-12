package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	LOG_LEVEL := os.Getenv("LOG_LEVEL")
	var logLevel zerolog.Level
	if LOG_LEVEL == "" {
		logLevel = zerolog.InfoLevel
	} else {
		var err error
		if logLevel, err = zerolog.ParseLevel(LOG_LEVEL); err != nil {
			logLevel = zerolog.InfoLevel
		}
	}

	// 设置全局时区为本地时区
	zerolog.TimeFieldFormat = time.RFC3339

	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		Logger = zerolog.New(os.Stderr).
			Level(logLevel).
			With().
			Timestamp(). // 添加时间戳
			Logger()
	} else {
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		Logger = zerolog.New(file).
			Level(logLevel).
			With().
			Timestamp(). // 添加时间戳
			Logger()
	}

	Logger.Info().Msg("Log level: " + logLevel.String())
}
