package controller

import (
	"app-helley/src/helper"
	usecase "app-helley/src/usecase/user"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController interface {
	Store(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	DeleteById(http.ResponseWriter, *http.Request)
	Find(http.ResponseWriter, *http.Request)
	FindById(http.ResponseWriter, *http.Request)
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

func (u *userController) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	var storeUser helper.StoreUserRequest
	if err = json.Unmarshal(body, &storeUser); err != nil {
		helper.GetError(w, err)
		return
	}

	response, err := u.storeUserUseCase.Handle(storeUser)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	if errEncode := helper.RequestEncode(w, response); errEncode != nil {
		w.WriteHeader(500)
	}
}

func (u *userController) Update(w http.ResponseWriter, r *http.Request) {
	helper.SetContentTypeJson(w)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	var userUpdate helper.UpdateUserRequest
	if err = json.Unmarshal(body, &userUpdate); err != nil {
		helper.GetError(w, err)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	response, err := u.updateUserUseCase.Handle(id, userUpdate)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	if errEncode := helper.RequestEncode(w, response); errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (u *userController) DeleteById(w http.ResponseWriter, r *http.Request) {
	helper.SetContentTypeJson(w)

	params := mux.Vars(r)
	id := params["id"]

	response, err := u.deleteUserUseCase.Handle(id)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	if errEncode := helper.RequestEncode(w, response); errEncode != nil {
		helper.GetError(w, err)
		return
	}
}

func (u *userController) Find(w http.ResponseWriter, r *http.Request) {
	helper.SetContentTypeJson(w)

	response, err := u.usersUseCase.Handle()

	if err != nil {
		helper.GetError(w, err)
		return
	}

	if errEncode := helper.RequestEncode(w, response); errEncode != nil {
		helper.GetError(w, err)
	}
}

func (u *userController) FindById(w http.ResponseWriter, r *http.Request) {
	helper.SetContentTypeJson(w)
	helper.SetContentTypeJson(w)

	params := mux.Vars(r)
	id := params["id"]

	response, err := u.userUseCase.Handle(id)

	if err != nil {
		helper.GetError(w, err)
		return
	}

	if errEncode := helper.RequestEncode(w, response); errEncode != nil {
		helper.GetError(w, err)
	}

}
