package setup

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Validator = &Validator{validator: validator.New()}
	e.HTTPErrorHandler = HTTPErrorHandler

	return e
}
