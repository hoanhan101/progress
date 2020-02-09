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

	if err := c.Validate(n); err != nil {
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
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	goal, err := model.GetGoal(h.db, params["id"])
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goal)
}

// UpdateGoal updates a specified goal in the system.
func (h *Handler) UpdateGoal(c echo.Context) error {
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	values, err := getFormValues(c, map[string]bool{
		"name": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	goal, err := model.UpdateGoal(h.db, params["id"], values["name"])
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, goal)
}

// DeleteGoal deletes a specified goal in the system.
func (h *Handler) DeleteGoal(c echo.Context) error {
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	err = model.DeleteGoal(h.db, params["id"])
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
