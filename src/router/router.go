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

	var routesHandle []routes.Route
	routesHandle = append(routesHandle, userRoutes...)

	for _, route := range routesHandle {
		router.HandleFunc(route.Path, route.HandleFunc).Methods(route.Method)
	}

	return router
}
