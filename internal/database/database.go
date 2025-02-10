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
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Name,
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
