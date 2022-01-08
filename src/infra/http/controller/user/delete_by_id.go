package controller_user

import (
	user "app-helley/src/app/usecase/user"
	"app-helley/src/infra/http/controller"
	"app-helley/src/infra/http/dto"
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
