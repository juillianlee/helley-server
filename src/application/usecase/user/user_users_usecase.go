package user

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/contract"
)

type UsersUseCase interface {
	Handle() ([]contract.UserResponse, error)
}

type usersUserCase struct {
	userRepository app_repository.UserRepository
}

func NewUsersUseCase(userRepository app_repository.UserRepository) UsersUseCase {
	return &usersUserCase{
		userRepository: userRepository,
	}
}

func (u *usersUserCase) Handle() ([]contract.UserResponse, error) {

	users, err := u.userRepository.Find()

	if err != nil {
		return []contract.UserResponse{}, err
	}

	var response []contract.UserResponse

	for _, user := range users {
		response = append(response, contract.UserResponse{
			ID:    user.ID.Hex(),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}
