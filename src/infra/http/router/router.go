package router

import (
	"app-helley/src/infra/config"
	"app-helley/src/infra/http/router/routes"
	"app-helley/src/infra/http/setup"
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

	for _, route := range routesHandle {
		switch route.Method {
		case http.MethodPut:
			e.PUT(route.Path, route.HandleFunc)
		case http.MethodGet:
			e.GET(route.Path, route.HandleFunc)
		case http.MethodPost:
			e.POST(route.Path, route.HandleFunc)
		case http.MethodDelete:
			e.DELETE(route.Path, route.HandleFunc)
		}
	}

	return e
}
