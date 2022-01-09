package controller

import (
	user "app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/controller"
	"net/http"

	"github.com/labstack/echo"
)

type findByIdUserController struct {
	usecase user.UserUseCase
}

func NewFindByIdUserController(usecase user.UserUseCase) controller.Handler {
	return &findByIdUserController{
		usecase: usecase,
	}
}

func (h *findByIdUserController) Handle(c echo.Context) error {
	id := c.Param("id")
	response, err := h.usecase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
