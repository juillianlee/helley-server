package user

import (
	app_repository "app-helley/src/app/repository"
)

type DeleteUserUseCase interface {
	Handle(id string) error
}

type deleteUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewDeleteUserUseCase(userRepository app_repository.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

func (u *deleteUserUseCase) Handle(id string) error {
	return u.userRepository.DeleteById(id)
}
