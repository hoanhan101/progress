package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Config(c echo.Context) error {
	return c.JSON(http.StatusOK, h.config)
}
