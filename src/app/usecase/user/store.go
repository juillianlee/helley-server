package user

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
)

type StoreUserModel struct {
	Name     string
	Email    string
	Password string
}

type StoreUserUseCase interface {
	Handle(storeUser StoreUserModel) (domain_user.User, error)
}

type storeUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewStoreUserUseCase(userRepository app_repository.UserRepository) StoreUserUseCase {
	return &storeUserUseCase{
		userRepository: userRepository,
	}
}

func (usecase *storeUserUseCase) Handle(storeUser StoreUserModel) (domain_user.User, error) {

	user := domain_user.User{
		Name:     storeUser.Name,
		Email:    storeUser.Email,
		Password: storeUser.Password,
	}

	if err := user.Validate(); err != nil {
		return user, nil
	}

	password, err := domain_user.GenerateHashPassword(storeUser.Password)
	if err != nil {
		return domain_user.User{}, err
	}

	user.Password = password

	return usecase.userRepository.Store(user)

}
