package controller

import (
	usecase_account "app-helley/src/app/usecase/account"
	"app-helley/src/app/validator"
	app_validator "app-helley/src/app/validator"
	"app-helley/src/infra/http/controller"
	"app-helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type loginController struct {
	usecase usecase_account.LoginUseCase
}

func NewLoginController(loginUseCase usecase_account.LoginUseCase) controller.Handler {
	return &loginController{
		usecase: loginUseCase,
	}
}

// Handler responsável por realizar o login no sistema via api
func (h *loginController) Handle(c echo.Context) (err error) {

	payload := new(dto.LoginRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if err := app_validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    validator.ValidationErrors(err).Error(),
		})
	}

	response, err := h.usecase.Handle(payload.Username, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
