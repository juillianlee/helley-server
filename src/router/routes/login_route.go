package routes

import (
	"app-helley/src/config"
	"app-helley/src/controller"
	"app-helley/src/usecase/login"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRoutes(db *mongo.Database) []config.Route {
	loginUseCase := login.NewLoginUseCase()
	handler := controller.NewLoginController(loginUseCase)

	return makeLoginRoutes(handler)
}

func makeLoginRoutes(handler controller.LoginController) []config.Route {
	return []config.Route{
		{
			Path:                   "/login",
			Method:                 http.MethodPost,
			HandleFunc:             handler.Login,
			RequiredAuthentication: false,
		},
	}
}
