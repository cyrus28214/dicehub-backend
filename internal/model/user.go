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
	Name      *string   `db:"name" json:"name"`
	Avatar    *string   `db:"avatar" json:"avatar"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserFromJson User

func (u *UserFromJson) Scan(value any) error {
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
        returning *
    `, openId)

	if err != nil {
		return nil, err
	}

	log.Logger.Debug().Msgf("user: %+v", user)

	return &user, nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	err := database.DB.Get(&user, `
        select *
        from "user"
        where id = $1
    `, id)

	if err != nil {
		log.Logger.Error().Err(err).Int64("id", id).Msg("Failed to get user by id")
		return nil, err
	}

	return &user, nil
}

// update name and avatar
func UpdateUser(userId int64, avatar string, name string) error {
	if avatar != "" {
		_, err := database.DB.Exec(`
			update "user"
			set avatar = $1
			where id = $2
		`, avatar, userId)
		if err != nil {
			return err
		}
	}

	if name != "" {
		_, err := database.DB.Exec(`
			update "user"
			set name = $1
			where id = $2
		`, name, userId)
		if err != nil {
			return err
		}
	}

	return nil
}
