package usecase

import (
	app_repository "helley/src/app/repository"
	app_security "helley/src/app/security"
)

type (
	RefreshTokenUseCase interface {
		Handle(refreshToken string) (app_security.TokenPayload, error)
	}

	refreshTokenUseCase struct {
		tokenManager   app_security.TokenManager
		userRepository app_repository.UserRepository
	}
)

func NewRefreshTokenUseCase(tokenManager app_security.TokenManager, userRepository app_repository.UserRepository) RefreshTokenUseCase {
	return &refreshTokenUseCase{
		tokenManager:   tokenManager,
		userRepository: userRepository,
	}
}

func (r *refreshTokenUseCase) Handle(refreshToken string) (app_security.TokenPayload, error) {
	claims, err := r.tokenManager.ValidateRefreshToken(refreshToken)

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	userId := claims["sub"].(string)

	user, err := r.userRepository.FindById(userId)

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	tokenPayload, err := r.tokenManager.GenerateTokenPair(user)

	if err != nil {
		return app_security.TokenPayload{}, err
	}

	return tokenPayload, err
}
