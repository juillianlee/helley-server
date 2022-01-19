package controller

import (
	user "helley/src/app/usecase/user"
	"helley/src/app/validator"
	"helley/src/infra/http/controller"
	"helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type storeUserController struct {
	usecase user.StoreUserUseCase
}

func NewStoreUserHandler(usecase user.StoreUserUseCase) controller.Handler {
	return &storeUserController{
		usecase: usecase,
	}
}

func (h *storeUserController) Handle(c echo.Context) error {
	storeUser := new(dto.StoreUserRequest)
	if err := c.Bind(storeUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(storeUser); err != nil {
		return err
	}

	response, err := h.usecase.Handle(user.StoreUserModel{
		Name:     storeUser.Name,
		Email:    storeUser.Email,
		Password: storeUser.Password,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}
