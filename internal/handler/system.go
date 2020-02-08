package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateSystem create a system in the system.
func (h *Handler) CreateSystem(c echo.Context) error {
	values, err := getFormValues(c, map[string]bool{
		"goal_id": true,
		"name":    true,
		"repeat":  false,
	})

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	systems, err := model.CreateSystem(h.db, values["goal_id"], values["name"], values["repeat"])
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, systems)
}

// GetSystems gets all systems in the system.
func (h *Handler) GetSystems(c echo.Context) error {
	systems, err := model.GetSystems(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, systems)
}
