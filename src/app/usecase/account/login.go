package usecase

import (
	app_repository "helley/src/app/repository"
	app_security "helley/src/app/security"
)

type LoginUseCase interface {
	Handle(username string, password string) (app_security.TokenPayload, error)
}

type loginUseCase struct {
	tokenManager   app_security.TokenManager
	userRepository app_repository.UserRepository
}

func NewLoginUseCase(tokenManager app_security.TokenManager, userRepository app_repository.UserRepository) LoginUseCase {
	return &loginUseCase{
		tokenManager:   tokenManager,
		userRepository: userRepository,
	}
}

/*
	Handler responsavel para validar o acesso do usuario
	e retornar os tokens de acessos.
*/
func (h *loginUseCase) Handle(username string, password string) (app_security.TokenPayload, error) {

	user, err := h.userRepository.FindByEmail(username)

	if err != nil {
		return app_security.TokenPayload{}, app_security.ErrUserNameOrPasswordInvalid
	}

	if err != nil || !user.CheckPasswordHash(password) {
		return app_security.TokenPayload{}, app_security.ErrUserNameOrPasswordInvalid
	}

	response, err := h.tokenManager.GenerateTokenPair(user)

	if err != nil {
		return app_security.TokenPayload{}, app_security.ErrUserNameOrPasswordInvalid
	}

	return response, nil

}
