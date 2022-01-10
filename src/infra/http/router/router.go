package router

import (
	"app-helley/src/infra/config"
	"app-helley/src/infra/http/router/routes"
	"app-helley/src/infra/http/setup"
	"app-helley/src/infra/middleware"
	repository "app-helley/src/infra/repository/mongo"
	"app-helley/src/infra/security"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
Generates new router to application with router configured
*/
func NewRouter(db *mongo.Database) *echo.Echo {
	e := setup.SetupRouter()

	userRoutes := routes.NewUserRoutes(db)
	loginRoutes := routes.NewLoginRoutes(db)
	websocketRoutes := routes.NewWebsocketRoutes()

	var routesHandle []config.Route
	routesHandle = append(routesHandle, userRoutes...)
	routesHandle = append(routesHandle, loginRoutes...)
	routesHandle = append(routesHandle, websocketRoutes...)

	userRepository := repository.NewUserRepository(db)
	tokenManager := security.NewTokenManager(config.ACESSS_TOKEN_SECRET, config.REFRESH_TOKEN_SECRET)
	authMiddleware := middleware.NewAuthMiddleware(tokenManager, userRepository)

	var globalMiddlewares []echo.MiddlewareFunc

	for _, route := range routesHandle {
		middlewares := globalMiddlewares

		if route.RequiredAuthentication {
			middlewares = append(middlewares, authMiddleware.Middleware)
		}

		switch route.Method {
		case http.MethodPut:
			e.PUT(route.Path, route.HandleFunc, middlewares...)
		case http.MethodGet:
			e.GET(route.Path, route.HandleFunc, middlewares...)
		case http.MethodPost:
			e.POST(route.Path, route.HandleFunc, middlewares...)
		case http.MethodDelete:
			e.DELETE(route.Path, route.HandleFunc, middlewares...)
		}
	}

	return e
}
