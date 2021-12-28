package usecase

import (
	"app-helley/src/helper"
	"app-helley/src/repository"
)

type UpdateUserUseCase interface {
	Handle(id string, helper *helper.UpdateUserRequest) (helper.MessageResponse, error)
}

type updateUseUseCase struct {
	userRepository repository.UserRepository
}

func NewUpdateUserUseCase(userRepository repository.UserRepository) UpdateUserUseCase {
	return &updateUseUseCase{
		userRepository: userRepository,
	}
}

func (usecase *updateUseUseCase) Handle(id string, h *helper.UpdateUserRequest) (helper.MessageResponse, error) {
	user, err := usecase.userRepository.FindById(id)

	if err != nil {
		return helper.MessageResponse{}, err
	}

	user.Name = h.Name
	user.Email = h.Email

	err = usecase.userRepository.Update(user)
	if err != nil {
		return helper.MessageResponse{}, err
	}

	return helper.MessageResponse{
		Message: "Usuário atualizado com sucesso",
	}, nil
}
