package router

import (
	"app-helley/src/router/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
Generates new router to application with router configured
*/
func NewRouter(db *mongo.Database) *mux.Router {
	router := mux.NewRouter()

	userRoutes := routes.CreateUserRoutes(db)
	chatRoutes := routes.CreateChatRoutes(db)

	var routesHandle []routes.Route
	routesHandle = append(routesHandle, userRoutes...)
	routesHandle = append(routesHandle, chatRoutes...)

	for _, route := range routesHandle {
		r := router.HandleFunc(route.Path, route.HandleFunc)
		if route.Method != "" {
			r.Methods(route.Method)
		}
	}

	return router
}
