package controller

import (
	"app-helley/src/helper"
	usecase "app-helley/src/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Store(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	DeleteById(c echo.Context) (err error)
	Find(c echo.Context) (err error)
	FindById(c echo.Context) (err error)
}

type userController struct {
	storeUserUseCase  usecase.StoreUserUseCase
	updateUserUseCase usecase.UpdateUserUseCase
	deleteUserUseCase usecase.DeleteUserUseCase
	usersUseCase      usecase.UsersUseCase
	userUseCase       usecase.UserUseCase
}

func NewUserController(
	storeUserUseCase usecase.StoreUserUseCase,
	updateUserUseCase usecase.UpdateUserUseCase,
	deleteUserUseCase usecase.DeleteUserUseCase,
	usersUseCase usecase.UsersUseCase,
	userUseCase usecase.UserUseCase,
) UserController {
	return &userController{
		storeUserUseCase:  storeUserUseCase,
		updateUserUseCase: updateUserUseCase,
		deleteUserUseCase: deleteUserUseCase,
		usersUseCase:      usersUseCase,
		userUseCase:       userUseCase,
	}
}

func (u *userController) Store(c echo.Context) (err error) {
	storeUser := new(helper.StoreUserRequest)
	if err := c.Bind(storeUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(storeUser); err != nil {
		return err
	}

	response, err := u.storeUserUseCase.Handle(storeUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *userController) Update(c echo.Context) (err error) {
	userUpdate := new(helper.UpdateUserRequest)

	if err := c.Bind(userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(userUpdate); err != nil {
		return err
	}

	id := c.Param("id")

	response, err := u.updateUserUseCase.Handle(id, userUpdate)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *userController) DeleteById(c echo.Context) (err error) {
	id := c.Param("id")

	response, err := u.deleteUserUseCase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *userController) Find(c echo.Context) (err error) {

	response, err := u.usersUseCase.Handle()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *userController) FindById(c echo.Context) (err error) {
	id := c.Param("id")
	response, err := u.userUseCase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
