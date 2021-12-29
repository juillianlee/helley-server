package routes

import (
	"app-helley/src/config"
	"app-helley/src/controller"
	"app-helley/src/repository"
	usecase "app-helley/src/usecase/user"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserRoutes(db *mongo.Database) []config.Route {
	userRepository := repository.NewUserRepository(db)
	storeUserUseCase := usecase.NewStoreUserUseCase(userRepository)
	userUpdateUseCase := usecase.NewUpdateUserUseCase(userRepository)
	deleteUserUseCase := usecase.NewDeleteUserUseCase(userRepository)
	usersUseCase := usecase.NewUsersUseCase(userRepository)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(
		storeUserUseCase,
		userUpdateUseCase,
		deleteUserUseCase,
		usersUseCase,
		userUseCase,
	)
	return makeUserRoutes(userController)
}

func makeUserRoutes(handler controller.UserController) []config.Route {
	return []config.Route{
		{
			Path:                   "/users",
			Method:                 http.MethodGet,
			HandleFunc:             handler.Find,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/users",
			Method:                 http.MethodPost,
			HandleFunc:             handler.Store,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodGet,
			HandleFunc:             handler.FindById,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodPut,
			HandleFunc:             handler.Update,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodDelete,
			HandleFunc:             handler.DeleteById,
			RequiredAuthentication: false,
		},
	}
}
