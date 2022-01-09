package routes

import (
	usecase_account "app-helley/src/app/usecase/account"
	"app-helley/src/infra/config"
	controller_account "app-helley/src/infra/http/controller/account"
	repository_mongo "app-helley/src/infra/repository/mongo"
	"app-helley/src/infra/security"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRoutes(db *mongo.Database) []config.Route {
	tokenManager := security.NewTokenManager(config.ACESSS_TOKEN_SECRET, config.REFRESH_TOKEN_SECRET)
	userRepository := repository_mongo.NewUserRepository(db)

	loginUseCase := usecase_account.NewLoginUseCase(tokenManager, userRepository)
	loginControllerHandler := controller_account.NewLoginController(loginUseCase)

	refreshTokenUseCase := usecase_account.NewRefreshTokenUseCase(tokenManager)
	refreshTokenControllerHandler := controller_account.NewRefreshTokenController(refreshTokenUseCase)

	return []config.Route{
		{
			Path:                   "/login",
			Method:                 http.MethodPost,
			HandleFunc:             loginControllerHandler.Handle,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/refreshToken",
			Method:                 http.MethodPost,
			HandleFunc:             refreshTokenControllerHandler.Handle,
			RequiredAuthentication: false,
		},
	}
}
