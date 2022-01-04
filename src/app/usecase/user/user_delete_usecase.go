package user

import (
	app_repository "app-helley/src/app/repository"
	"app-helley/src/presentation"
)

type DeleteUserUseCase interface {
	Handle(id string) (presentation.MessageResponse, error)
}

type deleteUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewDeleteUserUseCase(userRepository app_repository.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

func (u *deleteUserUseCase) Handle(id string) (presentation.MessageResponse, error) {
	if err := u.userRepository.DeleteById(id); err != nil {
		return presentation.MessageResponse{}, err
	}

	return presentation.MessageResponse{
		Message: "Usuário deletado com sucesso",
	}, nil
}
