package user

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/presentation"
)

type UsersUseCase interface {
	Handle() ([]presentation.UserResponse, error)
}

type usersUserCase struct {
	userRepository app_repository.UserRepository
}

func NewUsersUseCase(userRepository app_repository.UserRepository) UsersUseCase {
	return &usersUserCase{
		userRepository: userRepository,
	}
}

func (u *usersUserCase) Handle() ([]presentation.UserResponse, error) {

	users, err := u.userRepository.Find()

	if err != nil {
		return []presentation.UserResponse{}, err
	}

	var response []presentation.UserResponse

	for _, user := range users {
		response = append(response, presentation.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}
