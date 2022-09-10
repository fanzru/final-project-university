package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

var (
	ErrTokenInvalid = errors.New("token invalid")
)

func EncodeToken(id int64, email string, secret string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: id,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	})

	return token.SignedString([]byte(secret))
}

func DecodeToken(token string, secret string) (Claims, error) {
	claims := &Claims{}
	tokenDecode, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if !tokenDecode.Valid {
		return Claims{}, ErrTokenInvalid
	}
	return *claims, nil
}
