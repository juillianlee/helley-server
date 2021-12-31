package controller

import (
	"app-helley/src/application/usecase/login"
	"app-helley/src/contract"
	"fmt"
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
	payload := new(contract.LoginRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := controller.loginUseCase.Handle(payload.Username, payload.Password)
	if err != nil {
		fmt.Println("iasjsiodjioasda")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *loginController) RefreshToken(c echo.Context) (err error) {
	payload := new(contract.RefreshTokenRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := controller.refreshTokenUseCase.Handle(payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
