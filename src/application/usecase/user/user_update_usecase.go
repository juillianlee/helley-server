package user

import (
	"app-helley/src/contract"
	"app-helley/src/infrastructure/repository"
)

type UpdateUserUseCase interface {
	Handle(id string, contract *contract.UpdateUserRequest) (contract.MessageResponse, error)
}

type updateUseUseCase struct {
	userRepository repository.UserRepository
}

func NewUpdateUserUseCase(userRepository repository.UserRepository) UpdateUserUseCase {
	return &updateUseUseCase{
		userRepository: userRepository,
	}
}

func (usecase *updateUseUseCase) Handle(id string, h *contract.UpdateUserRequest) (contract.MessageResponse, error) {
	user, err := usecase.userRepository.FindById(id)

	if err != nil {
		return contract.MessageResponse{}, err
	}

	user.Name = h.Name
	user.Email = h.Email

	err = usecase.userRepository.Update(user)
	if err != nil {
		return contract.MessageResponse{}, err
	}

	return contract.MessageResponse{
		Message: "Usuário atualizado com sucesso",
	}, nil
}
