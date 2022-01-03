package controller

import (
	usecase "app-helley/src/application/usecase/user"
	"app-helley/src/presentation"
	"net/http"

	"github.com/labstack/echo"
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

// Create new user on database
func (u *userController) Store(c echo.Context) (err error) {
	storeUser := new(presentation.StoreUserRequest)
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

// Update user on database
func (u *userController) Update(c echo.Context) (err error) {
	userUpdate := new(presentation.UpdateUserRequest)

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

// Delete user on database
func (u *userController) DeleteById(c echo.Context) (err error) {
	id := c.Param("id")

	response, err := u.deleteUserUseCase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// Find user on database
func (u *userController) Find(c echo.Context) (err error) {

	response, err := u.usersUseCase.Handle()

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

// Find by id on database
func (u *userController) FindById(c echo.Context) (err error) {
	id := c.Param("id")
	response, err := u.userUseCase.Handle(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
