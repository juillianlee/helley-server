package user

import (
	"app-helley/src/contract"
	"app-helley/src/infrastructure/repository"
)

type UserUseCase interface {
	Handle(id string) (contract.UserResponse, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Handle(id string) (contract.UserResponse, error) {
	user, err := u.userRepository.FindById(id)

	if err != nil {
		return contract.UserResponse{}, err
	}

	return contract.UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
