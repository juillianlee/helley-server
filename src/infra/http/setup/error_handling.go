package setup

import (
	app_repository "app-helley/src/app/repository"
	"app-helley/src/app/security"
	"app-helley/src/infra/http/dto"
	"net/http"

	"errors"

	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {

	switch {
	case errors.Is(err, app_repository.ErrNotFoundRegister):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	case errors.Is(err, app_repository.ErrDuplicateKey):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	case errors.Is(err, security.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
	default:
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	}
}
