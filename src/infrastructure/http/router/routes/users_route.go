package routes

import (
	"app-helley/src/application/usecase/user"
	"app-helley/src/infrastructure/config"
	"app-helley/src/infrastructure/http/controller"
	repository_mongo "app-helley/src/infrastructure/repository/mongo"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRoutes(db *mongo.Database) []config.Route {
	userRepository := repository_mongo.NewUserRepository(db)
	storeUserUseCase := user.NewStoreUserUseCase(userRepository)
	userUpdateUseCase := user.NewUpdateUserUseCase(userRepository)
	deleteUserUseCase := user.NewDeleteUserUseCase(userRepository)
	usersUseCase := user.NewUsersUseCase(userRepository)
	userUseCase := user.NewUserUseCase(userRepository)
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
