package controller_account

import (
	"app-helley/src/app/usecase/login"
	"app-helley/src/infra/http/controller"
	"app-helley/src/presentation"
	"net/http"

	"github.com/labstack/echo"
)

type refreshTokenController struct {
	refreshTokenUseCase login.RefreshTokenUseCase
}

func NewRefreshTokenController(refreshTokenUseCase login.RefreshTokenUseCase) controller.Handler {
	return &refreshTokenController{
		refreshTokenUseCase: refreshTokenUseCase,
	}
}

func (h *refreshTokenController) Handle(c echo.Context) (err error) {
	payload := new(presentation.RefreshTokenRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.refreshTokenUseCase.Handle(payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
