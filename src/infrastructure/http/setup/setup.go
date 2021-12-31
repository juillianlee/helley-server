package setup

import (
	"app-helley/src/infrastructure/http/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
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

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Validator = &Validator{validator: validator.New()}
	e.HTTPErrorHandler = errors.ErrorHandler

	return e
}