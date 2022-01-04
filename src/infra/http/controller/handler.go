package controller

import "github.com/labstack/echo"

type Handler interface {
	Handle(c echo.Context) error
}
