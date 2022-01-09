package usecase

import (
	app_repository "app-helley/src/app/repository"
	app_security "app-helley/src/app/security"
	usecase "app-helley/src/app/usecase/user"
	repository_memory "app-helley/src/infra/repository/memory"
	"app-helley/src/infra/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStoreUser(userRepository *repository_memory.UserRepository) {
	store := usecase.NewStoreUserUseCase(userRepository)
	store.Handle(usecase.StoreUserModel{
		Name:     "Juillian Lee",
		Email:    "juillian.lee@gmail.com",
		Password: "abc123",
	})
}

func TestLoginUseCaseSuccessfuly(t *testing.T) {
	tokenManager := security.NewTokenManager("abc123")

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

	tokenManager := security.NewTokenManager("abc123")

	userRepository := &repository_memory.UserRepository{}

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@gmail.com", "abc123")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_repository.ErrNotFoundRegister)
}

func TestLoginUseCasePasswordError(t *testing.T) {
	tokenManager := security.NewTokenManager("abc123")

	userRepository := &repository_memory.UserRepository{}

	testStoreUser(userRepository)

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@gmail.com", "abc")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_security.ErrUnauthorized)
}

func TestLoginUseCaseUsernameInvalid(t *testing.T) {
	tokenManager := security.NewTokenManager("abc123")

	userRepository := &repository_memory.UserRepository{}

	testStoreUser(userRepository)

	loginUsecase := NewLoginUseCase(tokenManager, userRepository)

	response, err := loginUsecase.Handle("juillian.lee@hotmail.com", "abc123")

	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
	assert.ErrorIs(t, err, app_repository.ErrNotFoundRegister)
}
