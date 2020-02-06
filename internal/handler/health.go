package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// Health is the handler for `/health` endpoint.
func (h *Handler) Health(c echo.Context) error {
	health, err := model.GetHealth(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, health)
}
