package routes

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Path                   string
	Method                 string
	HandleFunc             func(c echo.Context) (err error)
	RequiredAuthentication bool
}
