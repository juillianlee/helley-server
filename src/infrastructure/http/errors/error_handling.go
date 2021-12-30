package errors

import (
	"app-helley/src/contract"

	"errors"

	"github.com/labstack/echo"
)

func ErrorHandler(err error, c echo.Context) {
	err = WrapError(err)
	var (
		msg    string = "Internal server error"
		status int    = 500
		apiErr APIError
	)
	if errors.As(err, &apiErr) {
		status, msg = apiErr.APIError()
	}
	c.JSON(status, contract.ErrorResponse{Message: msg, StatusCode: status})
}
