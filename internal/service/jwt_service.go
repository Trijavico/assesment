package service

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

var UnAuthorized = errors.New("unauthorized")

type JwtService struct {
	sync.Once
	secret string
}

func (jwtService *JwtService) CreateToken(claims jwt.MapClaims) (string, error) {
	jwtService.Do(func() {
		jwtService.secret = os.Getenv("JWT_SECRET")
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtService.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *JwtService) ParseToken(tokenString string) (jwt.MapClaims, error) {
	jwtService.Do(func() {
		jwtService.secret = os.Getenv("JWT_SECRET")
	})

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.secret), nil
	})

	if err != nil {
		return nil, UnAuthorized
	}
	if !token.Valid {
		return nil, UnAuthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, UnAuthorized
	}

	return claims, nil
}
