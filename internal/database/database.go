package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"wx-miniprogram-backend/internal/config"
	"wx-miniprogram-backend/internal/log"
)

var DB *sqlx.DB

func Connect() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DB.Host,
		config.Cfg.DB.Port,
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.Name,
	)

	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Logger.Fatal().Msgf("failed to connect to database: %v", err)
	}

	log.Logger.Info().Msg("connected to database")
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Logger.Error().Msgf("failed to close database: %v", err)
	}
}
