package login

// TODO remover a infra daqui para ficar somente cadama de dominio aqui.
import (
	app_security "app-helley/src/app/security"
	"errors"
)

type LoginUseCase interface {
	Handle(username string, password string) (map[string]string, error)
}

type loginUseCase struct {
	tokenService app_security.TokenManager
}

func NewLoginUseCase(tokenService app_security.TokenManager) LoginUseCase {
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
