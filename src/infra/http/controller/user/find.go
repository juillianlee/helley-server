package controller

import (
	user "app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/controller"
	"net/http"

	"github.com/labstack/echo"
)

type findUsersController struct {
	usecase user.UsersUseCase
}

func NewFindUsersController(usecase user.UsersUseCase) controller.Handler {
	return &findUsersController{
		usecase: usecase,
	}
}

func (h *findUsersController) Handle(c echo.Context) error {
	response, err := h.usecase.Handle()

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
