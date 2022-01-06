package login

import (
	app_security "app-helley/src/app/security"
)

type (
	RefreshTokenUseCase interface {
		Handle(refreshToken string) (app_security.TokenPayload, error)
	}

	refreshTokenUseCase struct {
		tokenManager app_security.TokenManager
	}
)

func NewRefreshTokenUseCase(tokenManager app_security.TokenManager) RefreshTokenUseCase {
	return &refreshTokenUseCase{
		tokenManager: tokenManager,
	}
}

func (r *refreshTokenUseCase) Handle(refreshToken string) (app_security.TokenPayload, error) {
	return r.tokenManager.RefreshToken(refreshToken)
}
