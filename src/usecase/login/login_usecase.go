package login

import (
	"app-helley/src/service"
	"errors"
)

type LoginUseCase interface {
	Handle(username string, password string) (map[string]string, error)
}

type loginUseCase struct {
	tokenService service.TokenService
}

func NewLoginUseCase(tokenService service.TokenService) LoginUseCase {
	return &loginUseCase{
		tokenService: tokenService,
	}
}

func (h *loginUseCase) Handle(username string, password string) (map[string]string, error) {

	if username == "juillian" && password == "abc123" {
		response, err := h.tokenService.GenerateTokenPair()

		if err != nil {
			return map[string]string{}, err
		}

		return response, nil

	}

	return nil, errors.New("ErrUnauthorized")
}
