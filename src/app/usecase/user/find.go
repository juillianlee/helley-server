package user

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
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
