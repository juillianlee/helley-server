package user

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
	"app-helley/src/presentation"
)

type StoreUserUseCase interface {
	Handle(storeUser *presentation.StoreUserRequest) (presentation.StoreUserResponse, error)
}

type storeUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewStoreUserUseCase(userRepository app_repository.UserRepository) StoreUserUseCase {
	return &storeUserUseCase{
		userRepository: userRepository,
	}
}

func (usecase *storeUserUseCase) Handle(storeUser *presentation.StoreUserRequest) (presentation.StoreUserResponse, error) {
	user := domain_user.User{
		Name:     storeUser.Name,
		Email:    storeUser.Email,
		Password: storeUser.Password,
	}

	userInserted, err := usecase.userRepository.Store(user)

	if err != nil {
		return presentation.StoreUserResponse{}, err
	}

	response := presentation.StoreUserResponse{
		ID:    userInserted.ID,
		Name:  userInserted.Name,
		Email: user.Email,
	}

	return response, nil
}
