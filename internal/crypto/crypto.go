package crypto

import (
	"fmt"
	"time"
	"wx-miniprogram-backend/internal/config"
	"wx-miniprogram-backend/internal/log"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

func SignToken(Id int64) (string, error) {
	// Token 过期时间设置为30天
	expiresAt := time.Now().Add(30 * 24 * time.Hour)

	claims := Claims{
		Id: Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.Cfg.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return config.Cfg.JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	err = fmt.Errorf("invalid token")
	log.Logger.Error().Str("token", tokenString).Err(err).Msg("Failed to verify token")
	return nil, err
}
