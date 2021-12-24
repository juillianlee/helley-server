package usecase

import (
	"app-helley/src/helper"
	"app-helley/src/repository"
)

type UserUseCase interface {
	Handle(id string) (helper.UserResponse, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Handle(id string) (helper.UserResponse, error) {
	user, err := u.userRepository.FindById(id)

	if err != nil {
		return helper.UserResponse{}, err
	}

	return helper.UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
