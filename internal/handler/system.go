package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// ListSystems is the handler for listing all systems endpoint.
func (h *Handler) ListSystems(c echo.Context) error {
	systems, err := model.ListSystems(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, systems)
}
