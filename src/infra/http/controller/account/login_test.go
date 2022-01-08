package controller_account

import (
	app_security "app-helley/src/app/security"
	"app-helley/src/app/usecase/login"
	usecase "app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/setup"
	repository_memory "app-helley/src/infra/repository/memory"
	"app-helley/src/infra/security"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
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

// Realiza o teste do cenario de api de um login realizado com sucesso
func TestLoginSucessfuly(t *testing.T) {
	e := setup.SetupRouter()

	userRepository := &repository_memory.UserRepository{}
	testStoreUser(userRepository)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username": "juillian.lee@gmail.com", "password": "abc123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	tokenManager := security.NewTokenManager("secret")

	loginUseCase := login.NewLoginUseCase(tokenManager, userRepository)
	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
	}
}

// Realiza o teste do cenario de um login que vai dar falha
func TestLoginFail(t *testing.T) {
	e := setup.SetupRouter()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username": "not found", "password": ""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tokenManager := security.NewTokenManager("secret")
	loginUseCase := login.NewLoginUseCase(tokenManager, &repository_memory.UserRepository{})

	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Empty(t, result.AccessToken)
		assert.Empty(t, result.RefreshToken)
	}
}

// Realiza o teste de um envio de dados invalidos.
func TestLoginEmptyBody(t *testing.T) {
	e := setup.SetupRouter()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(``))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tokenManager := security.NewTokenManager("secret")
	loginUseCase := login.NewLoginUseCase(tokenManager, &repository_memory.UserRepository{})
	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Empty(t, result.AccessToken)
		assert.Empty(t, result.RefreshToken)
	}
}
