package login

import (
	app_security "app-helley/src/app/security"
)

type LoginUseCase interface {
	Handle(username string, password string) (app_security.TokenPayload, error)
}

type loginUseCase struct {
	tokenManager app_security.TokenManager
}

func NewLoginUseCase(tokenService app_security.TokenManager) LoginUseCase {
	return &loginUseCase{
		tokenManager: tokenService,
	}
}

func (h *loginUseCase) Handle(username string, password string) (app_security.TokenPayload, error) {

	if username == "juillian" && password == "abc123" {
		response, err := h.tokenManager.GenerateTokenPair()

		if err != nil {
			return app_security.TokenPayload{}, err
		}

		return response, nil

	}

	return app_security.TokenPayload{}, app_security.ErrUnauthorized
}
