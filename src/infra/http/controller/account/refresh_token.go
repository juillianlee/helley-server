package controller_account

import (
	"app-helley/src/app/usecase/login"
	app_validator "app-helley/src/app/validator"
	"app-helley/src/infra/http/controller"
	"app-helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type refreshTokenController struct {
	refreshTokenUseCase login.RefreshTokenUseCase
}

// Retorna uma instancia do handler para que seja possivel fazer o refresh token
func NewRefreshTokenController(refreshTokenUseCase login.RefreshTokenUseCase) controller.Handler {
	return &refreshTokenController{
		refreshTokenUseCase: refreshTokenUseCase,
	}
}

// Realiza o refresh token do usuario
func (h *refreshTokenController) Handle(c echo.Context) (err error) {
	payload := new(dto.RefreshTokenRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := app_validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, app_validator.ValidationErrors(err))
	}

	response, err := h.refreshTokenUseCase.Handle(payload.RefreshToken)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
