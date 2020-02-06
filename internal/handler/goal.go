package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateGoal is the handler creating a goal endpoint.
func (h *Handler) CreateGoal(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"require name value",
		)
	}

	g, err := model.CreateGoal(h.db, name)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, g)
}

// ListGoals is the handler for listing all goals endpoint.
func (h *Handler) ListGoals(c echo.Context) error {
	goals, err := model.ListGoals(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, goals)
}
