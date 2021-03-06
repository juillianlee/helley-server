package controller

import (
	user "helley/src/app/usecase/user"
	"helley/src/infra/http/controller"
	"helley/src/infra/http/dto"
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

	err := h.usecase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "Usuário deletado com sucesso",
	})
}
