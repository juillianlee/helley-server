package login

import (
	app_repository "app-helley/src/app/repository"
	app_security "app-helley/src/app/security"
	domain_user "app-helley/src/domain/user"
)

type LoginUseCase interface {
	Handle(username string, password string) (app_security.TokenPayload, error)
}

type loginUseCase struct {
	tokenManager   app_security.TokenManager
	userRepository app_repository.UserRepository
}

func NewLoginUseCase(tokenService app_security.TokenManager, userRepository app_repository.UserRepository) LoginUseCase {
	return &loginUseCase{
		tokenManager:   tokenService,
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
		return app_security.TokenPayload{}, err
	}

	hashPassword, err := domain_user.GenerateHashPassword(password)

	if err != nil || !user.CheckPasswordHash(hashPassword) {
		return app_security.TokenPayload{}, app_security.ErrUnauthorized
	}

	response, err := h.tokenManager.GenerateTokenPair()

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	return response, nil

}
