package controller

import (
	"encoding/json"
	usecase "helley/src/app/usecase/account"
	"helley/src/infra/http/dto"
	"helley/src/infra/http/setup"
	repository "helley/src/infra/repository/memory"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountSuccessfully(t *testing.T) {

	userRepository := &repository.UserRepository{}

	usecase := usecase.NewCreateAccountUseCase(userRepository)

	e := setup.SetupRouter()
	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(`{"name": "Juillian Lee", "email": "juillian.lee@gmail.com", "password": "acb123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := NewCreateAccountController(usecase).Handle(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response dto.UserResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.Name)
	assert.NotEmpty(t, response.Email)
}

func TestCreateAccountBadRequestEmptyBody(t *testing.T) {
	userRepository := &repository.UserRepository{}

	usecase := usecase.NewCreateAccountUseCase(userRepository)

	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(`{}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := setup.SetupRouter()
	c := e.NewContext(req, rec)

	err := NewCreateAccountController(usecase).Handle(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response dto.ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, response.StatusCode, http.StatusBadRequest)
	assert.NotEmpty(t, response.Message)
}
