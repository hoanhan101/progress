package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Config is the handler for `/config` endpoint.
func (h *Handler) Config(c echo.Context) error {
	return c.JSON(http.StatusOK, h.config)
}
