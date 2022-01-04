package controller_account

import (
	"app-helley/src/app/usecase/login"
	"app-helley/src/infra/http/controller"
	"app-helley/src/presentation"
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
	payload := new(presentation.LoginRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.usecase.Handle(payload.Username, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
