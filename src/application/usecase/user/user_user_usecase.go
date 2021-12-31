package user

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/contract"
)

type UserUseCase interface {
	Handle(id string) (contract.UserResponse, error)
}

type userUseCase struct {
	userRepository app_repository.UserRepository
}

func NewUserUseCase(userRepository app_repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Handle(id string) (contract.UserResponse, error) {
	user, err := u.userRepository.FindById(id)

	if err != nil {
		return contract.UserResponse{}, err
	}

	return contract.UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
