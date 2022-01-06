package setup

import (
	"github.com/labstack/echo"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = HTTPErrorHandler

	return e
}
