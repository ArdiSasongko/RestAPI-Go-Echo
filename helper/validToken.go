package helper

import (
	"errors"
	"strings"
)

func ValidToken(token string) (string, error) {
	if token == "" || !strings.HasPrefix(token, "Bearer") {
		return "", errors.New("authHeader need")
	}

	validToken := strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	if validToken == "" {
		return "", errors.New("need token")
	}

	return validToken, nil
}
