package usecase

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain"
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

func (usecase *updateUseUseCase) Handle(id string, u UpdateUserModel) (domain_user.User, error) {
	user, err := usecase.userRepository.FindById(id)

	if err != nil {
		return domain_user.User{}, err
	}

	user.Name = u.Name
	user.Email = u.Email

	if err = user.Validate(); err != nil {
		return user, nil
	}

	err = usecase.userRepository.Update(user)

	return user, err
}
