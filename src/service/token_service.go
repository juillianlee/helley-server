package service

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	TokenService interface {
		GenerateTokenPair() (map[string]string, error)
		RefreshToken() (string, error)
		AccessToken() (string, error)
	}

	tokenService struct {
		jwtSecret string
	}
)

func NewTokenService(jwtSecret string) TokenService {
	return &tokenService{
		jwtSecret: jwtSecret,
	}
}

func (s *tokenService) AccessToken() (string, error) {
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

func (s *tokenService) RefreshToken() (string, error) {
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

func (s *tokenService) GenerateTokenPair() (map[string]string, error) {
	accessToken, err := s.AccessToken()

	if err != nil {
		return map[string]string{}, err
	}

	refreshToken, err := s.RefreshToken()
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
