/*
	Caso de uso para tratar da criação do usuário
	em caso de sucesso deve, adicionar o usuário para ser enviado
	um e-mail de onboard e retornar os dados do usuário de dominio
	inserido.
*/

package usecase

import (
	"app-helley/src/app/repository"
	"app-helley/src/app/validator"
	"app-helley/src/domain"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type CreateAccountModel struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type CreateAccountUseCase interface {
	Handle(CreateAccountModel) (domain.User, error)
}

type createAccountUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateAccountUseCase(userRepository repository.UserRepository) CreateAccountUseCase {
	return &createAccountUseCase{
		userRepository: userRepository,
	}
}

func (u *createAccountUseCase) Handle(createAccount CreateAccountModel) (domain.User, error) {
	if err := validator.Validate(createAccount); err != nil {
		log.Infof("Invalid data on createAccount: %v", validator.ValidationErrors(err))
		return domain.User{}, validator.ValidationErrors(err)
	}

	user, err := u.userRepository.FindByEmail(createAccount.Email)
	if err != nil && !errors.Is(err, repository.ErrNotFoundRegister) {
		log.Errorf("Transaction Find user by e-mail failed: %v", err)
		return domain.User{}, err
	}

	if user.ID != "" {
		log.Warningf("User by e-mail %s already exists", createAccount.Email)
		return domain.User{}, fmt.Errorf("user by e-mail %s already exists", createAccount.Email)
	}

	hashPassword, err := domain.GenerateHashPassword(createAccount.Password)
	if err != nil {
		log.Errorf("Generate hash password error: %v", err)
		return domain.User{}, err
	}

	newUser := domain.User{
		Name:      createAccount.Name,
		Email:     createAccount.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	user, err = u.userRepository.Store(newUser)

	if err != nil {
		log.Errorf("Transaction store user failed: %v", err)
		return domain.User{}, err
	}

	log.Infof("User %s successfully registered e-mail %s", user.ID, user.Email)

	// TODO adicionar o usuário a fila de envio de e-mail do onboard

	return user, nil

}
