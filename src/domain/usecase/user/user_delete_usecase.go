package user

import (
	"app-helley/src/helper"
	"app-helley/src/repository"
)

type DeleteUserUseCase interface {
	Handle(id string) (helper.MessageResponse, error)
}

type deleteUserUseCase struct {
	userRepository repository.UserRepository
}

func NewDeleteUserUseCase(userRepository repository.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

func (u *deleteUserUseCase) Handle(id string) (helper.MessageResponse, error) {
	if err := u.userRepository.DeleteById(id); err != nil {
		return helper.MessageResponse{}, err
	}

	return helper.MessageResponse{
		Message: "Usuário deletado com sucesso",
	}, nil
}