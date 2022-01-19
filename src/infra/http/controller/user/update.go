package controller

import (
	user "helley/src/app/usecase/user"
	"helley/src/infra/http/controller"
	"helley/src/infra/http/dto"
	"net/http"

	"github.com/labstack/echo"
)

type updateUserController struct {
	usecase user.UpdateUserUseCase
}

func NewUpdateUserController(usecase user.UpdateUserUseCase) controller.Handler {
	return &updateUserController{
		usecase: usecase,
	}
}

func (h *updateUserController) Handle(c echo.Context) error {
	userUpdate := new(dto.UpdateUserRequest)

	if err := c.Bind(userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(userUpdate); err != nil {
		return err
	}

	id := c.Param("id")

	response, err := h.usecase.Handle(id, user.UpdateUserModel{
		Name:  userUpdate.Name,
		Email: userUpdate.Email,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
