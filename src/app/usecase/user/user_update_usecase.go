package user

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
)

type UpdateUserModel struct {
	Name  string
	Email string
}

type UpdateUserUseCase interface {
	Handle(id string, presentation UpdateUserModel) (domain_user.User, error)
}

type updateUseUseCase struct {
	userRepository app_repository.UserRepository
}

func NewUpdateUserUseCase(userRepository app_repository.UserRepository) UpdateUserUseCase {
	return &updateUseUseCase{
		userRepository: userRepository,
	}
}

func (usecase *updateUseUseCase) Handle(id string, h UpdateUserModel) (domain_user.User, error) {
	user, err := usecase.userRepository.FindById(id)

	if err != nil {
		return domain_user.User{}, err
	}

	user.Name = h.Name
	user.Email = h.Email

	err = usecase.userRepository.Update(user)

	return user, err
}
