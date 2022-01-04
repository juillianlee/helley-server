package controller_user

import (
	"app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/controller"
	"app-helley/src/presentation"
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
	storeUser := new(presentation.StoreUserRequest)
	if err := c.Bind(storeUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(storeUser); err != nil {
		return err
	}

	response, err := h.usecase.Handle(storeUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}
