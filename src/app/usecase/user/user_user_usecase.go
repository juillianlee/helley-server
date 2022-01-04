package user

import (
	app_repository "app-helley/src/app/repository"
	"app-helley/src/presentation"
)

type UserUseCase interface {
	Handle(id string) (presentation.UserResponse, error)
}

type userUseCase struct {
	userRepository app_repository.UserRepository
}

func NewUserUseCase(userRepository app_repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Handle(id string) (presentation.UserResponse, error) {
	user, err := u.userRepository.FindById(id)

	if err != nil {
		return presentation.UserResponse{}, err
	}

	return presentation.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
