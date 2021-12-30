package login

import (
	"app-helley/src/helper"
	"app-helley/src/infrastructure/security"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type (
	RefreshTokenUseCase interface {
		Handle(refreshToken *helper.RefreshTokenRequest) (map[string]string, error)
	}

	refreshTokenUseCase struct {
		tokenService security.TokenManager
	}
)

func NewRefreshTokenUseCase(tokenService security.TokenManager) RefreshTokenUseCase {
	return &refreshTokenUseCase{
		tokenService: tokenService,
	}
}

func (r *refreshTokenUseCase) Handle(refreshToken *helper.RefreshTokenRequest) (map[string]string, error) {

	token, err := jwt.Parse(refreshToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int(claims["sub"].(float64)) == 1 {
			return r.tokenService.GenerateTokenPair()
		}
	}

	return map[string]string{}, err

}
