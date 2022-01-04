package controller_user

import (
	"app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/controller"
	"net/http"

	"github.com/labstack/echo"
)

type deleteUserByIdController struct {
	usecase user.DeleteUserUseCase
}

func NewDeleteUserByIdController(usecase user.DeleteUserUseCase) controller.Handler {
	return &deleteUserByIdController{
		usecase: usecase,
	}
}

func (h *deleteUserByIdController) Handle(c echo.Context) error {
	id := c.Param("id")

	response, err := h.usecase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
