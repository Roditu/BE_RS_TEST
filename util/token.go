package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
    secretKey string
}

func NewJWTMaker(key string) *JWTMaker {
    return &JWTMaker{secretKey: key}
}

func (maker *JWTMaker) CreateToken(userID int32, duration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(duration).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(maker.secretKey), nil
    })
}
