package routes

import "net/http"

type Route struct {
	Path                   string
	Method                 string
	HandleFunc             func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}
