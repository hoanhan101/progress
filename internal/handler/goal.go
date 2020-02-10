package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateGoal creates a goal in the system.
func (h *Handler) CreateGoal(c echo.Context) error {
	n := new(model.NewGoal)
	if err := c.Bind(n); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(n); err != nil {
		return errBadRequest(err)
	}

	goal, err := model.CreateGoal(h.db, n)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goal)
}

// GetGoals gets all goals in the system.
func (h *Handler) GetGoals(c echo.Context) error {
	goals, err := model.GetGoals(h.db)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goals)
}

// GetGoal gets a specified goal in the system.
func (h *Handler) GetGoal(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	goal, err := model.GetGoal(h.db, id)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goal)
}

// UpdateGoal updates a specified goal in the system.
func (h *Handler) UpdateGoal(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	u := new(model.UpdatedGoal)
	if err := c.Bind(u); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(u); err != nil {
		return errBadRequest(err)
	}

	goal, err := model.UpdateGoal(h.db, id, u)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goal)
}

// DeleteGoal deletes a specified goal in the system.
func (h *Handler) DeleteGoal(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	err := model.DeleteGoal(h.db, id)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "deleted successfully",
		},
	)
}
