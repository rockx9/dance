package service

import (
	"dance/types"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = "d1f5e4e376b243c190e62b6b4e4db8d8"
var ErrInvalidToken = errors.New("invalid token")

type IAuth interface {
	Verify(username, password string) (string, error)
	ParseToken(tokenstring string) (*types.AuthUser, error)
}

type Auth struct {
}

func NewAuth() IAuth {
	return &Auth{}
}

func (a Auth) Verify(username, password string) (string, error) {
	if username != "admin" || password != "abcd.1234" {
		return "", errors.New("Incorrect username or password.")
	}
	claims := types.AuthUser{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}

func (a Auth) ParseToken(tokenstring string) (*types.AuthUser, error) {
	if tokenstring == "" {
		return nil, ErrInvalidToken
	}
	token, err := jwt.ParseWithClaims(tokenstring, &types.AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(*types.AuthUser); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken

}
