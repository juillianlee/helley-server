package usecase

import (
	"app-helley/src/helper"
	"app-helley/src/repository"
)

type UsersUseCase interface {
	Handle() ([]helper.UserResponse, error)
}

type usersUserCase struct {
	userRepository repository.UserRepository
}

func NewUsersUseCase(userRepository repository.UserRepository) UsersUseCase {
	return &usersUserCase{
		userRepository: userRepository,
	}
}

func (u *usersUserCase) Handle() ([]helper.UserResponse, error) {

	users, err := u.userRepository.Find()

	if err != nil {
		return []helper.UserResponse{}, err
	}

	var response []helper.UserResponse

	for _, user := range users {
		response = append(response, helper.UserResponse{
			ID:    user.ID.Hex(),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}
