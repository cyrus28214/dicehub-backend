package model

import (
	"database/sql"
	"time"

	"wx-miniprogram-backend/internal/database"
)

type Comment struct {
	Id        int64        `db:"id" json:"id"`
	UserId    int64        `db:"user_id" json:"user_id"`
	GameId    int64        `db:"game_id" json:"game_id"`
	Content   string       `db:"content" json:"content"`
	Rating    float64      `db:"rating" json:"rating"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	User      UserFromJson `db:"user" json:"user,omitempty"`
}

// CreateComment 创建评论
func CreateComment(userId, gameId int64, content string, rating float64) error {
	_, err := database.DB.Exec(`
		insert into "comment" (user_id, game_id, content, rating)
		values ($1, $2, $3, $4)
	`, userId, gameId, content, rating)
	return err
}

// GetGameComments 获取游戏的评论列表
func GetGameComments(gameId int64) ([]Comment, error) {
	comments := []Comment{}
	query := `
		select 
			c.*,
			row_to_json(u) as user
		from "comment" c
		join "user" u on c.user_id = u.id
		where c.game_id = $1
		order by c.created_at desc
	`
	err := database.DB.Select(&comments, query, gameId)
	return comments, err
}

// UpdateComment 更新评论
func UpdateComment(id, userId int64, content string, rating float64) error {
	result, err := database.DB.Exec(`
		update "comment"
		set content = $1, rating = $2, updated_at = now()
		where id = $3 and user_id = $4
	`, content, rating, id, userId)

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

// DeleteComment 删除评论
func DeleteComment(id, userId int64) error {
	result, err := database.DB.Exec(`
		delete from "comment"
		where id = $1 and user_id = $2
	`, id, userId)

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

// GetUserGameComment 获取用户对特定游戏的评论
func GetUserGameComment(userId, gameId int64) (*Comment, error) {
	var comment Comment
	err := database.DB.Get(&comment, `
		select *
		from "comment"
		where user_id = $1 and game_id = $2
	`, userId, gameId)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
