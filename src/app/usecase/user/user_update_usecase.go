package user

import (
	app_repository "app-helley/src/app/repository"
	"app-helley/src/presentation"
)

type UpdateUserUseCase interface {
	Handle(id string, presentation *presentation.UpdateUserRequest) (presentation.MessageResponse, error)
}

type updateUseUseCase struct {
	userRepository app_repository.UserRepository
}

func NewUpdateUserUseCase(userRepository app_repository.UserRepository) UpdateUserUseCase {
	return &updateUseUseCase{
		userRepository: userRepository,
	}
}

func (usecase *updateUseUseCase) Handle(id string, h *presentation.UpdateUserRequest) (presentation.MessageResponse, error) {
	user, err := usecase.userRepository.FindById(id)

	if err != nil {
		return presentation.MessageResponse{}, err
	}

	user.Name = h.Name
	user.Email = h.Email

	err = usecase.userRepository.Update(user)
	if err != nil {
		return presentation.MessageResponse{}, err
	}

	return presentation.MessageResponse{
		Message: "Usuário atualizado com sucesso",
	}, nil
}
