package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func errBadRequest(err error) error {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		err.Error(),
	)
}

func errInternalServer(err error) error {
	return echo.NewHTTPError(
		http.StatusInternalServerError,
		err.Error(),
	)
}
