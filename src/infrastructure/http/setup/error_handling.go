package setup

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/presentation"
	"net/http"

	"errors"

	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {

	switch {
	case errors.Is(err, app_repository.ErrNoResults):
		c.JSON(http.StatusNotFound, presentation.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	case errors.Is(err, app_repository.ErrDuplicateKey):
		c.JSON(http.StatusBadRequest, presentation.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	default:
		c.JSON(http.StatusInternalServerError, presentation.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	}
}
