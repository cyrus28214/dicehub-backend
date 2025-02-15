package model

import (
	"time"

	"wx-miniprogram-backend/internal/database"

	"github.com/lib/pq"
)

type Game struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Image       string    `db:"image" json:"image"`
	Rating      float64   `db:"rating" json:"rating"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	LikesCount  int       `db:"likes_count" json:"likes_count"`
	ExtraInfo   *string   `db:"extra_info" json:"extra_info"`
}

type GamDetail struct {
	Game
	Tags  TagArray `db:"tags" json:"tags"`
	Liked bool     `db:"liked" json:"liked"`
}

// GetGames 获取游戏列表，支持标签过滤
// 用户没登录就是id = 0，此时liked = false
func GetGames(tagIds pq.Int64Array, userId int64) ([]GamDetail, error) {
	var games []GamDetail
	var query string

	if len(tagIds) == 0 {
		query = `
            SELECT 
                g.*,
                json_agg(t) as tags,
                CASE WHEN count(l.user_id) > 0 THEN true ELSE false END as liked
            FROM game g
            LEFT JOIN game_tag_relation gt ON g.id = gt.game_id
            LEFT JOIN tag t ON gt.tag_id = t.id
            LEFT JOIN "like" l ON g.id = l.game_id AND l.user_id = $1
            GROUP BY g.id
        `
		err := database.DB.Select(&games, query, userId)
		if err != nil {
			return nil, err
		}

		return games, nil
	}

	// 如果指定了标签，通过标签过滤游戏
	query = `
        SELECT 
            g.*,
            json_agg(t) as tags,
            CASE WHEN count(l.user_id) > 0 THEN true ELSE false END as liked
        FROM game g
        LEFT JOIN game_tag_relation gt ON g.id = gt.game_id
        LEFT JOIN tag t ON gt.tag_id = t.id
        LEFT JOIN "like" l ON g.id = l.game_id AND l.user_id = $1
        GROUP BY g.id
        HAVING array_agg(t.id) @> $2
    `
	err := database.DB.Select(&games, query, userId, tagIds)
	if err != nil {
		return nil, err
	}

	return games, nil
}

// GetGameById 根据ID获取游戏详情
func GetGameById(id int64, userId int64) (*GamDetail, error) {
	query := `
        SELECT 
            g.*,
            json_agg(t) as tags,
            CASE WHEN count(l.user_id) > 0 THEN true ELSE false END as liked
        FROM game g
        LEFT JOIN game_tag_relation gt ON g.id = gt.game_id
        LEFT JOIN tag t ON gt.tag_id = t.id
        LEFT JOIN "like" l ON g.id = l.game_id AND l.user_id = $2
        WHERE g.id = $1
        GROUP BY g.id
    `

	var game GamDetail
	err := database.DB.Get(&game, query, id, userId)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
