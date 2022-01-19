package usecase

import (
	app_repository "helley/src/app/repository"
	domain_user "helley/src/domain"
)

type UsersUseCase interface {
	Handle() ([]domain_user.User, error)
}

type usersUserCase struct {
	userRepository app_repository.UserRepository
}

func NewUsersUseCase(userRepository app_repository.UserRepository) UsersUseCase {
	return &usersUserCase{
		userRepository: userRepository,
	}
}

func (u *usersUserCase) Handle() ([]domain_user.User, error) {
	users, err := u.userRepository.Find()
	return users, err
}
