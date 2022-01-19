package controller

import (
	usecase_account "helley/src/app/usecase/account"
	app_validator "helley/src/app/validator"
	"helley/src/infra/http/controller"
	"helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type refreshTokenController struct {
	refreshTokenUseCase usecase_account.RefreshTokenUseCase
}

// Retorna uma instancia do handler para que seja possivel fazer o refresh token
func NewRefreshTokenController(refreshTokenUseCase usecase_account.RefreshTokenUseCase) controller.Handler {
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
