package router

import (
	"app-helley/src/config"
	"app-helley/src/router/routes"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

/**
Generates new router to application with router configured
*/
func NewRouter(db *mongo.Database) *echo.Echo {
	e := echo.New()

	e.Validator = &Validator{validator: validator.New()}

	userRoutes := routes.CreateUserRoutes(db)
	websocketRoutes := routes.CreateWebsocketRoutes()

	var routesHandle []config.Route
	routesHandle = append(routesHandle, userRoutes...)
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
