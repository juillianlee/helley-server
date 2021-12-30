package user

import (
	"app-helley/src/contract"
	"app-helley/src/infrastructure/repository"
)

type DeleteUserUseCase interface {
	Handle(id string) (contract.MessageResponse, error)
}

type deleteUserUseCase struct {
	userRepository repository.UserRepository
}

func NewDeleteUserUseCase(userRepository repository.UserRepository) DeleteUserUseCase {
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
