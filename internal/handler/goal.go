package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateGoal creates a goal in the system.
func (h *Handler) CreateGoal(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"name is not specified in the request body",
		)
	}

	goal, err := model.CreateGoal(h.db, name)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, goal)
}

// GetGoals gets all goals in the system.
func (h *Handler) GetGoals(c echo.Context) error {
	goals, err := model.GetGoals(h.db)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, goals)
}

// GetGoal gets a specified goal in the system.
func (h *Handler) GetGoal(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"id is not specified in the URI",
		)
	}

	goal, err := model.GetGoal(h.db, id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, goal)
}

// UpdateGoal updates a specified goal in the system.
func (h *Handler) UpdateGoal(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"id is not specified in the URI",
		)
	}

	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"name is not specified in the request body",
		)
	}

	goal, err := model.UpdateGoal(h.db, id, name)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, goal)
}

// DeleteGoal deletes a specified goal in the system.
func (h *Handler) DeleteGoal(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"id is not specified in the URI",
		)
	}

	err := model.DeleteGoal(h.db, id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(http.StatusNoContent, nil)
}
