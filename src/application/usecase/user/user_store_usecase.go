package user

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/contract"
	"app-helley/src/domain"
)

type StoreUserUseCase interface {
	Handle(storeUser *contract.StoreUserRequest) (contract.StoreUserResponse, error)
}

type storeUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewStoreUserUseCase(userRepository app_repository.UserRepository) StoreUserUseCase {
	return &storeUserUseCase{
		userRepository: userRepository,
	}
}

func (usecase *storeUserUseCase) Handle(storeUser *contract.StoreUserRequest) (contract.StoreUserResponse, error) {
	user := domain.User{
		Name:     storeUser.Name,
		Email:    storeUser.Email,
		Password: storeUser.Password,
	}

	userInserted, err := usecase.userRepository.Store(user)

	if err != nil {
		return contract.StoreUserResponse{}, err
	}

	response := contract.StoreUserResponse{
		ID:    userInserted.ID.Hex(),
		Name:  userInserted.Name,
		Email: user.Email,
	}

	return response, nil
}
