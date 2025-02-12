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
}

type GameWithTags struct {
	Game
	Tags TagArray `db:"tags" json:"tags"`
}

// GetGames 获取游戏列表，支持标签过滤
func GetGames(tagIds pq.Int64Array) ([]GameWithTags, error) {
	var games []GameWithTags
	var query string

	if len(tagIds) == 0 {
		// 如果没有指定标签，获取所有游戏
		query = `
            select 
                g.*,
                json_agg(t) as tags
            from game g
            left join game_tag_relation gt on g.id = gt.game_id
            left join tag t on gt.tag_id = t.id
            group by g.id
        `
		err := database.DB.Select(&games, query)
		if err != nil {
			return nil, err
		}

		return games, nil
	}

	// 如果指定了标签，通过标签过滤游戏
	query = `
		SELECT 
			g.*,
			json_agg(t) as tags
		FROM game g
		LEFT JOIN game_tag_relation gt ON g.id = gt.game_id
		LEFT JOIN tag t ON gt.tag_id = t.id
		GROUP BY g.id
		HAVING array_agg(t.id) @> $1;
    `
	err := database.DB.Select(&games, query, tagIds)
	if err != nil {
		return nil, err
	}

	return games, nil

}

// GetGameById 根据ID获取游戏详情
func GetGameById(id int64) (*GameWithTags, error) {
	query := `
        select 
            g.*,
            json_agg(t) as tags
        from game g
        left join game_tag_relation gt on g.id = gt.game_id
        left join tag t on gt.tag_id = t.id
        where g.id = $1
        group by g.id
    `

	var game GameWithTags
	err := database.DB.Get(&game, query, id)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
