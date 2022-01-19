package security

import (
	app_security "helley/src/app/security"
	"helley/src/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	tokenManager struct {
		keyAccessToken  string
		keyRefreshToken string
	}
)

func NewTokenManager(keyAcessToken string, keyRefreshToken string) app_security.TokenManager {
	return &tokenManager{
		keyAccessToken:  keyAcessToken,
		keyRefreshToken: keyRefreshToken,
	}
}

func (s *tokenManager) GenerateTokenPair(u domain.User) (app_security.TokenPayload, error) {
	accessToken, err := s.accessToken(u)

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	refreshToken, err := s.refreshToken(u)
	if err != nil {
		return app_security.TokenPayload{}, err
	}

	return app_security.TokenPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *tokenManager) accessToken(u domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = u.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	accessToken, err := token.SignedString([]byte(s.keyAccessToken))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *tokenManager) refreshToken(u domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshToken, err := token.SignedString([]byte(s.keyRefreshToken))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *tokenManager) ValidateRefreshToken(refreshToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, app_security.ErrUnexpectedSignin
		}

		return []byte(s.keyRefreshToken), nil
	})

	if err != nil {
		return map[string]interface{}{}, app_security.ErrUnexpectedSignin
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return map[string]interface{}{}, app_security.ErrUnexpectedSignin
}

func (s *tokenManager) ValidateAccessToken(accessToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, app_security.ErrUnexpectedSignin
		}

		return []byte(s.keyAccessToken), nil
	})

	if err != nil {
		return map[string]interface{}{}, app_security.ErrUnexpectedSignin
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return map[string]interface{}{}, app_security.ErrUnexpectedSignin
}
