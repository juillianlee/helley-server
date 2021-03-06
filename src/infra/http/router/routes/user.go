package routes

import (
	user "helley/src/app/usecase/user"
	"helley/src/infra/config"
	controller_user "helley/src/infra/http/controller/user"
	repository_mongo "helley/src/infra/repository/mongo"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRoutes(db *mongo.Database) []config.Route {
	userRepository := repository_mongo.NewUserRepository(db)

	deleteUserByIdHandler := controller_user.NewDeleteUserByIdController(user.NewDeleteUserUseCase(userRepository))
	findByIdUserHandler := controller_user.NewFindByIdUserController(user.NewUserUseCase(userRepository))
	findUsersHandler := controller_user.NewFindUsersController(user.NewUsersUseCase(userRepository))
	updateUserHandler := controller_user.NewUpdateUserController(user.NewUpdateUserUseCase(userRepository))
	storeUserHandler := controller_user.NewStoreUserHandler(user.NewStoreUserUseCase(userRepository))

	return []config.Route{
		{
			Path:                   "/users",
			Method:                 http.MethodGet,
			HandleFunc:             findUsersHandler.Handle,
			RequiredAuthentication: true,
		},
		{
			Path:                   "/users",
			Method:                 http.MethodPost,
			HandleFunc:             storeUserHandler.Handle,
			RequiredAuthentication: true,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodGet,
			HandleFunc:             findByIdUserHandler.Handle,
			RequiredAuthentication: true,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodPut,
			HandleFunc:             updateUserHandler.Handle,
			RequiredAuthentication: true,
		},
		{
			Path:                   "/users/:id",
			Method:                 http.MethodDelete,
			HandleFunc:             deleteUserByIdHandler.Handle,
			RequiredAuthentication: true,
		},
	}
}
