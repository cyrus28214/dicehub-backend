package model

import (
	"encoding/json"
	"errors"
	"time"
	"wx-miniprogram-backend/internal/database"
	"wx-miniprogram-backend/internal/log"
)

type Tag struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Image       *string   `db:"image" json:"image"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type TagArray []Tag

func (t *TagArray) Scan(value any) error {
	if value == nil {
		*t = []Tag{}
		return nil
	}

	// 检查是否是字节切片
	bytes, ok := value.([]byte)
	if !ok {
		log.Logger.Error().Msgf("failed to unmarshal Tags value: %v", value)
		return errors.New("failed to unmarshal Tags value")
	}

	log.Logger.Debug().Str("bytes", string(bytes)).Msg("tags")

	str := string(bytes)

	if str == "" || str == "[null]" {
		*t = []Tag{}
		return nil
	}

	// 解析 JSON 字符串
	var tags []Tag
	if err := json.Unmarshal([]byte(str), &tags); err != nil {
		log.Logger.Debug().Str("str", str).Msg("tags")
		log.Logger.Error().Err(err).Msgf("failed to unmarshal Tags: %v", err)
		return err
	}

	*t = tags
	return nil
}

func GetTags() ([]Tag, error) {
	var tags []Tag
	err := database.DB.Select(&tags, "SELECT * FROM tag")
	return tags, err
}
