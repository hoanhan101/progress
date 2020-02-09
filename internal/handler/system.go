package handler

import (
	"errors"
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
		return errBadRequest(err)
	}

	systems, err := model.CreateSystem(h.db, values["goal_id"], values["name"], values["repeat"])
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, systems)
}

// GetSystems gets all systems in the system.
func (h *Handler) GetSystems(c echo.Context) error {
	systems, err := model.GetSystems(h.db)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, systems)
}

// GetSystem gets a specified system in the system.
func (h *Handler) GetSystem(c echo.Context) error {
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	system, err := model.GetSystem(h.db, params["id"])
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, system)
}

// UpdateSystem updates a specified system in the system.
func (h *Handler) UpdateSystem(c echo.Context) error {
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	values, err := getFormValues(c, map[string]bool{
		"name":   false,
		"repeat": false,
	})
	if err != nil {
		return errBadRequest(err)
	}

	if values["name"] == "" && values["repeat"] == "" {
		return errBadRequest(errors.New("either name or repeat value must be specified in the request body"))
	}

	system, err := model.UpdateSystem(h.db, params["id"], values["name"], values["repeat"])
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, system)
}

// DeleteSystem deletes a specified system in the system.
func (h *Handler) DeleteSystem(c echo.Context) error {
	params, err := getParams(c, map[string]bool{
		"id": true,
	})
	if err != nil {
		return errBadRequest(err)
	}

	err = model.DeleteSystem(h.db, params["id"])
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
