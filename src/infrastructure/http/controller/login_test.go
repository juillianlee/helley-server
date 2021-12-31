package controller

import (
	"app-helley/src/application/usecase/login"
	"app-helley/src/infrastructure/http/setup"
	"app-helley/src/infrastructure/security"
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
	refreshTokenUseCase := login.NewRefreshTokenUseCase(tokenManager)
	handler := NewLoginController(loginUseCase, refreshTokenUseCase)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Contains(t, result, "access_token")
		assert.Contains(t, result, "refresh_token")
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
	refreshTokenUseCase := login.NewRefreshTokenUseCase(tokenManager)
	handler := NewLoginController(loginUseCase, refreshTokenUseCase)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.NotContains(t, result, "access_token")
		assert.NotContains(t, result, "refresh_token")
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
	refreshTokenUseCase := login.NewRefreshTokenUseCase(tokenManager)
	handler := NewLoginController(loginUseCase, refreshTokenUseCase)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)
		assert.NotContains(t, result, "access_token")
		assert.NotContains(t, result, "refresh_token")
	}
}
