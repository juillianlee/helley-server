package usecase

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain"
)

type UserUseCase interface {
	Handle(id string) (domain_user.User, error)
}

type userUseCase struct {
	userRepository app_repository.UserRepository
}

func NewUserUseCase(userRepository app_repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Handle(id string) (domain_user.User, error) {
	user, err := u.userRepository.FindById(id)
	return user, err

}
