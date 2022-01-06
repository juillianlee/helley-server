package controller_account

import (
	app_security "app-helley/src/app/security"
	"app-helley/src/app/usecase/login"
	"app-helley/src/infra/http/setup"
	"app-helley/src/infra/security"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLoginSucessfuly(t *testing.T) {
	e := setup.SetupRouter()

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username": "juillian", "password": "abc123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	tokenManager := security.NewTokenManager("secret")
	loginUseCase := login.NewLoginUseCase(tokenManager)
	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
	}
}

func TestLoginFail(t *testing.T) {
	e := setup.SetupRouter()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username": "not found", "password": ""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tokenManager := security.NewTokenManager("secret")
	loginUseCase := login.NewLoginUseCase(tokenManager)
	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Empty(t, result.AccessToken)
		assert.Empty(t, result.RefreshToken)
	}
}

func TestLoginEmptyBody(t *testing.T) {
	e := setup.SetupRouter()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(``))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	tokenManager := security.NewTokenManager("secret")
	loginUseCase := login.NewLoginUseCase(tokenManager)
	handler := NewLoginController(loginUseCase)

	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result app_security.TokenPayload
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Empty(t, result.AccessToken)
		assert.Empty(t, result.RefreshToken)
	}
}