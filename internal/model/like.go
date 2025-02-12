package model

import (
	"time"

	"wx-miniprogram-backend/internal/database"
)

type Like struct {
	UserId    int64     `db:"user_id" json:"user_id"`
	GameId    int64     `db:"game_id" json:"game_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// LikeGame 点赞游戏
func LikeGame(userId, gameId int64) error {

	_, err := database.DB.Exec(`
		insert into "like" (user_id, game_id)
		values ($1, $2)
		on conflict (user_id, game_id) do nothing
	`, userId, gameId)
	if err != nil {
		return err
	}

	return nil
}

// UnlikeGame 取消点赞
func UnlikeGame(userId, gameId int64) error {
	_, err := database.DB.Exec(`
		delete from "like"
		where user_id = $1 and game_id = $2
	`, userId, gameId)
	if err != nil {
		return err
	}

	return nil
}

// IsGameLiked 检查用户是否已点赞
func IsGameLiked(userId, gameId int64) (bool, error) {
	var exists bool
	err := database.DB.Get(&exists, `
		select exists(
			select 1
			from "like"
			where user_id = $1 and game_id = $2
		)
	`, userId, gameId)
	return exists, err
}
