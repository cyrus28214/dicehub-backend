package model

import (
	"time"

	"wx-miniprogram-backend/internal/database"
	"wx-miniprogram-backend/internal/log"
)

type User struct {
	Id        int64     `db:"id" json:"id"`
	OpenId    string    `db:"openid" json:"openid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// FindOrCreateByOpenId 通过openid查找用户，如果不存在则创建
func FindOrCreateByOpenId(openId string) (*User, error) {
	var user User

	err := database.DB.Get(&user, `
		INSERT INTO "user" (openid)
		VALUES ($1)
		ON CONFLICT (openid) DO UPDATE
		  SET updated_at = NOW()
		RETURNING id, openid, created_at, updated_at
	`, openId)

	if err != nil {
		log.Logger.Error().Err(err).Str("openid", openId).Msg("Failed to find or create user")
		return nil, err
	}

	log.Logger.Debug().Msgf("user: %+v", user)

	return &user, nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	err := database.DB.Get(&user, `
		SELECT id, openid, created_at, updated_at
		FROM "user"
		WHERE id = $1
	`, id)

	if err != nil {
		log.Logger.Error().Err(err).Int64("id", id).Msg("Failed to get user by id")
		return nil, err
	}

	return &user, nil
}
