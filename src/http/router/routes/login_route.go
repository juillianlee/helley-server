package routes

import (
	"app-helley/src/config"
	"app-helley/src/http/controller"
	"app-helley/src/service"
	"app-helley/src/usecase/login"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRoutes(db *mongo.Database) []config.Route {
	tokenService := service.NewTokenService(config.JWT_SECRET)

	loginUseCase := login.NewLoginUseCase(tokenService)
	refreshTokenUseCase := login.NewRefreshTokenUseCase(tokenService)
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
