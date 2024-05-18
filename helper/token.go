package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCaseInterface interface {
	GeneratedToken(claims CustomClaims) (string, error)
	DecodeToken(tokenString string) (*jwt.Token, error)
}

type TokenUseCaseImpl struct{}

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GeneratedToken(claims CustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, errToken := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (t *TokenUseCaseImpl) DecodeToken(tokenString string) (*jwt.Token, error) {
	token, errToken := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if errToken != nil {
		return nil, errToken
	}

	return token, nil
}
