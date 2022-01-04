package config

import (
	"github.com/labstack/echo"
)

type Route struct {
	Path                   string
	Method                 string
	HandleFunc             func(c echo.Context) (err error)
	RequiredAuthentication bool
}
