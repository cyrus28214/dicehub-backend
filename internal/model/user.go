package model

import (
	"encoding/json"
	"errors"
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

func (u *User) Scan(value any) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("value is not a byte slice")
	}

	return json.Unmarshal(bytes, u)
}

// FindOrCreateByOpenId 通过openid查找用户，如果不存在则创建
func FindOrCreateByOpenId(openId string) (*User, error) {
	var user User

	err := database.DB.Get(&user, `
        insert into "user" (openid)
        values ($1)
        on conflict (openid) do update
          set updated_at = now()
        returning id, openid, created_at, updated_at
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
        select id, openid, created_at, updated_at
        from "user"
        where id = $1
    `, id)

	if err != nil {
		log.Logger.Error().Err(err).Int64("id", id).Msg("Failed to get user by id")
		return nil, err
	}

	return &user, nil
}
