package security

import (
	app_security "app-helley/src/app/security"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	tokenManager struct {
		jwtSecret string
	}
)

func NewTokenManager(jwtSecret string) app_security.TokenManager {
	return &tokenManager{
		jwtSecret: jwtSecret,
	}
}

func (s *tokenManager) GenerateTokenPair() (map[string]string, error) {
	accessToken, err := s.accessToken()

	if err != nil {
		return map[string]string{}, err
	}

	refreshToken, err := s.refreshToken()
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (s *tokenManager) accessToken() (string, error) {
	jwtRefreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := jwtRefreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshToken, err := jwtRefreshToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *tokenManager) refreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = "Juillian"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	accessToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
