package weixin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wx-miniprogram-backend/internal/config"
	"wx-miniprogram-backend/internal/log"
)

type Code2SessionResponse struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}

func Code2Session(code string) (*Code2SessionResponse, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.Cfg.Weixin.AppId,
		config.Cfg.Weixin.AppSecret,
		code,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to request code2session")
		return nil, err
	}
	defer resp.Body.Close()

	var result Code2SessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Logger.Error().Err(err).Msg("Failed to decode code2session response")
		return nil, err
	}

	if result.ErrCode != 0 {
		log.Logger.Error().
			Int("errcode", result.ErrCode).
			Str("errmsg", result.ErrMsg).
			Msg("Code2Session failed")
		return nil, fmt.Errorf("code2session failed: %s", result.ErrMsg)
	}

	return &result, nil
}
