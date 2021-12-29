package user

import (
	"app-helley/src/domain"
	"app-helley/src/helper"
	"app-helley/src/repository"
)

type StoreUserUseCase interface {
	Handle(storeUser *helper.StoreUserRequest) (helper.StoreUserResponse, error)
}

type storeUserUseCase struct {
	userRepository repository.UserRepository
}

func NewStoreUserUseCase(userRepository repository.UserRepository) StoreUserUseCase {
	return &storeUserUseCase{
		userRepository: userRepository,
	}
}

func (usecase *storeUserUseCase) Handle(storeUser *helper.StoreUserRequest) (helper.StoreUserResponse, error) {
	user := domain.User{
		Name:     storeUser.Name,
		Email:    storeUser.Email,
		Password: storeUser.Password,
	}

	userInserted, err := usecase.userRepository.Store(user)

	if err != nil {
		return helper.StoreUserResponse{}, err
	}

	response := helper.StoreUserResponse{
		ID:    userInserted.ID.Hex(),
		Name:  userInserted.Name,
		Email: user.Email,
	}

	return response, err
}
