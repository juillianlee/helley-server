package controller

import (
	"app-helley/src/helper"
	"app-helley/src/usecase/login"
	"net/http"

	"github.com/labstack/echo"
)

type (
	LoginController interface {
		Login(c echo.Context) (err error)
		RefreshToken(c echo.Context) (err error)
	}

	loginController struct {
		loginUseCase        login.LoginUseCase
		refreshTokenUseCase login.RefreshTokenUseCase
	}
)

func NewLoginController(loginUseCase login.LoginUseCase, refreshTokenUseCase login.RefreshTokenUseCase) LoginController {
	return &loginController{
		loginUseCase: loginUseCase,
	}
}

func (controller *loginController) Login(c echo.Context) (err error) {
	payload := new(helper.LoginRequest)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := controller.loginUseCase.Handle(payload.Username, payload.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *loginController) RefreshToken(c echo.Context) (err error) {
	payload := new(helper.RefreshTokenRequest)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := controller.refreshTokenUseCase.Handle(payload)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
