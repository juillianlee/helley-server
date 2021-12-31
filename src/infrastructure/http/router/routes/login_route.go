package routes

import (
	"app-helley/src/application/usecase/login"
	"app-helley/src/infrastructure/config"
	"app-helley/src/infrastructure/http/controller"
	"app-helley/src/infrastructure/security"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRoutes(db *mongo.Database) []config.Route {
	tokenManager := security.NewTokenManager(config.JWT_SECRET)

	loginUseCase := login.NewLoginUseCase(tokenManager)
	refreshTokenUseCase := login.NewRefreshTokenUseCase(tokenManager)
	handler := controller.NewLoginController(loginUseCase, refreshTokenUseCase)

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
		{
			Path:                   "/refreshToken",
			Method:                 http.MethodPost,
			HandleFunc:             handler.Login,
			RequiredAuthentication: false,
		},
	}
}
