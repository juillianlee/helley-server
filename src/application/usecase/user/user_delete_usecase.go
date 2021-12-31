package user

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/contract"
)

type DeleteUserUseCase interface {
	Handle(id string) (contract.MessageResponse, error)
}

type deleteUserUseCase struct {
	userRepository app_repository.UserRepository
}

func NewDeleteUserUseCase(userRepository app_repository.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

func (u *deleteUserUseCase) Handle(id string) (contract.MessageResponse, error) {
	if err := u.userRepository.DeleteById(id); err != nil {
		return contract.MessageResponse{}, err
	}

	return contract.MessageResponse{
		Message: "Usuário deletado com sucesso",
	}, nil
}
