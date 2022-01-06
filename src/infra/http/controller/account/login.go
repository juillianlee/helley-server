package controller_account

import (
	"app-helley/src/app/usecase/login"
	"app-helley/src/infra/http/controller"
	"app-helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type loginController struct {
	usecase login.LoginUseCase
}

func NewLoginController(loginUseCase login.LoginUseCase) controller.Handler {
	return &loginController{
		usecase: loginUseCase,
	}
}

func (h *loginController) Handle(c echo.Context) (err error) {
	payload := new(dto.LoginRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.usecase.Handle(payload.Username, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		})
	}

	return c.JSON(http.StatusOK, response)
}
