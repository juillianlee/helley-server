package usecase

import (
	app_security "app-helley/src/app/security"
	repository_memory "app-helley/src/infra/repository/memory"
	"app-helley/src/infra/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStoreUser(userRepository *repository_memory.UserRepository) {
	NewCreateAccountUseCase(userRepository).Handle(CreateAccountModel{
		Name:     "Juillian Lee",
		Email:    "juillian.lee@gmail.com",
		Password: "abc123",
	})
}

func TestLoginUseCaseSuccessfuly(t *testing.T) {
	tokenManager := security.NewTokenManager("accessTokenSecret", "refreshTokenSecret")

	userRepository := &repository_memory.UserRepository{}

	testStoreUser(userRepository)

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@gmail.com", "abc123")

	if assert.NoError(t, err) {
		assert.NotEmpty(t, response.AccessToken)
		assert.NotEmpty(t, response.RefreshToken)
	}
}

func TestLoginUseCaseUserNotFound(t *testing.T) {

	tokenManager := security.NewTokenManager("accessTokenSecret", "refreshTokenSecret")

	userRepository := &repository_memory.UserRepository{}

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@gmail.com", "abc123")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_security.ErrUserNameOrPasswordInvalid)
}

func TestLoginUseCasePasswordError(t *testing.T) {
	tokenManager := security.NewTokenManager("accessTokenSecret", "refreshTokenSecret")

	userRepository := &repository_memory.UserRepository{}

	testStoreUser(userRepository)

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@gmail.com", "abc")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_security.ErrUserNameOrPasswordInvalid)
}

func TestLoginUseCaseUsernameInvalid(t *testing.T) {
	tokenManager := security.NewTokenManager("accessTokenSecret", "refreshTokenSecret")

	userRepository := &repository_memory.UserRepository{}

	testStoreUser(userRepository)

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@hotmail.com", "abc123")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_security.ErrUserNameOrPasswordInvalid)
}
