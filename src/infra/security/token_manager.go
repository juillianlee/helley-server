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

func (s *tokenManager) GenerateTokenPair() (app_security.TokenPayload, error) {
	accessToken, err := s.accessToken()

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	refreshToken, err := s.refreshToken()
	if err != nil {
		return app_security.TokenPayload{}, err
	}

	return app_security.TokenPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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

func (s *tokenManager) RefreshToken(refreshToken string) (app_security.TokenPayload, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, app_security.ErrUnexpectedSignin
		}

		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int(claims["sub"].(float64)) == 1 {
			return s.GenerateTokenPair()
		}
	}

	return app_security.TokenPayload{}, err
}
