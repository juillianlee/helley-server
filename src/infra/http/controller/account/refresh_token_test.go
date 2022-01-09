package controller

import (
	app_security "app-helley/src/app/security"
	usecase "app-helley/src/app/usecase/account"
	repository_memory "app-helley/src/infra/repository/memory"
	"app-helley/src/infra/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenSuccessfuly(t *testing.T) {

	userRepository := &repository_memory.UserRepository{}
	testStoreUser(userRepository)

	tokenManager := security.NewTokenManager("secretAccessToken", "secretRefreshToken")

	response, err := usecase.NewLoginUseCase(tokenManager, userRepository).Handle("juillian.lee@gmail.com", "abc123")
	assert.NoError(t, err)

	response, err = usecase.NewRefreshTokenUseCase(tokenManager).Handle(response.RefreshToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.AccessToken)
	assert.NotEmpty(t, response.RefreshToken)
}

func TestRefreshTokenInvalid(t *testing.T) {
	userRepository := &repository_memory.UserRepository{}
	testStoreUser(userRepository)

	tokenManager := security.NewTokenManager("secretAccessToken", "secretRefreshToken")

	response, err := usecase.NewLoginUseCase(tokenManager, userRepository).Handle("juillian.lee@gmail.com", "abc123")
	assert.NoError(t, err)

	response, err = usecase.NewRefreshTokenUseCase(tokenManager).Handle(response.AccessToken)
	assert.Error(t, err)
	assert.ErrorIs(t, err, app_security.ErrUnexpectedSignin)
	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
}

func TestRefreshTokenInvalidSubscription(t *testing.T) {
	userRepository := &repository_memory.UserRepository{}
	testStoreUser(userRepository)

	tokenManager := security.NewTokenManager("secretAccessToken", "secretRefreshToken")

	response, err := usecase.NewLoginUseCase(tokenManager, userRepository).Handle("juillian.lee@gmail.com", "abc123")
	assert.NoError(t, err)

	tokenManager = security.NewTokenManager("invalid access token", "Invalid refresh token")

	response, err = usecase.NewRefreshTokenUseCase(tokenManager).Handle(response.RefreshToken)
	assert.Error(t, err)
	assert.ErrorIs(t, err, app_security.ErrUnexpectedSignin)
	assert.Empty(t, response.AccessToken)
	assert.Empty(t, response.RefreshToken)
}
