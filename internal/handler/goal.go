package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateGoal creates a goal in the system.
func (h *Handler) CreateGoal(c echo.Context) error {
	n := c.FormValue("name")
	if n == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"name is not specified",
		)
	}

	g, err := model.CreateGoal(h.db, n)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, g)
}

// GetGoals gets all goals in the system.
func (h *Handler) GetGoals(c echo.Context) error {
	gs, err := model.GetGoals(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, gs)
}

// GetGoal gets a specified goal in the system.
func (h *Handler) GetGoal(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"id is not specified",
		)
	}

	g, err := model.GetGoal(h.db, id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, g)
}
