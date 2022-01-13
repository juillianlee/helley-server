package usecase

import (
	repository "app-helley/src/infra/repository/memory"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCaseSuccessfuly(t *testing.T) {

	userRepository := &repository.UserRepository{}

	user, err := NewCreateAccountUseCase(userRepository).Handle(CreateAccountModel{
		Name:     "Juillian Lee",
		Email:    "juillian.lee@gmail.com",
		Password: "abc123",
	})

	assert.NoError(t, err)
	assert.Equal(t, user.Name, "Juillian Lee")
	assert.Equal(t, user.Email, "juillian.lee@gmail.com")
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdateAt)
}

func TestCreateAccountUseCaseEmailAlreadyExists(t *testing.T) {
	userRepository := &repository.UserRepository{}

	_, err := NewCreateAccountUseCase(userRepository).Handle(CreateAccountModel{
		Name:     "Juillian Lee",
		Email:    "juillian.lee@gmail.com",
		Password: "abc123",
	})

	assert.NoError(t, err)

	_, err = NewCreateAccountUseCase(userRepository).Handle(CreateAccountModel{
		Name:     "Juillian Lee",
		Email:    "juillian.lee@gmail.com",
		Password: "abc123",
	})

	assert.Error(t, err)
	assert.Equal(t, err.Error(), fmt.Errorf("user by e-mail %s already exists", "juillian.lee@gmail.com").Error())
}

func TestCreateAccountUseCaseInvalidData(t *testing.T) {
	userRepository := &repository.UserRepository{}

	_, err := NewCreateAccountUseCase(userRepository).Handle(CreateAccountModel{})

	assert.Error(t, err)
}
