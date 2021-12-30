package user

import (
	"app-helley/src/contract"
	"app-helley/src/infrastructure/repository"
)

type UsersUseCase interface {
	Handle() ([]contract.UserResponse, error)
}

type usersUserCase struct {
	userRepository repository.UserRepository
}

func NewUsersUseCase(userRepository repository.UserRepository) UsersUseCase {
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
