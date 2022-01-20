package controller

import (
	usecase "helley/src/app/usecase/account"
	"helley/src/app/validator"
	"helley/src/infra/http/controller"
	"helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type createAccountController struct {
	usecase usecase.CreateAccountUseCase
}

func NewCreateAccountController(usecase usecase.CreateAccountUseCase) controller.Handler {
	return &createAccountController{
		usecase: usecase,
	}
}

func (h *createAccountController) Handle(c echo.Context) error {
	payload := new(dto.CreateAccountRequest)

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if err := validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := h.usecase.Handle(usecase.CreateAccountModel{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
