package log

import (
	"os"

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

	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		Logger = zerolog.New(os.Stderr).Level(logLevel)
	} else {
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		Logger = zerolog.New(file).Level(logLevel)
	}

	Logger.Info().Msg("Log level: " + logLevel.String())
}
