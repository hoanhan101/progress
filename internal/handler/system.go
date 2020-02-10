package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateSystem create a system in the system.
func (h *Handler) CreateSystem(c echo.Context) error {
	n := new(model.NewSystem)
	if err := c.Bind(n); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(n); err != nil {
		return errBadRequest(err)
	}

	system, err := model.CreateSystem(h.db, n)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, system)
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
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	system, err := model.GetSystem(h.db, id)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, system)
}

// UpdateSystem updates a specified system in the system.
func (h *Handler) UpdateSystem(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	u := new(model.UpdatedSystem)
	if err := c.Bind(u); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(u); err != nil {
		return errBadRequest(err)
	}

	system, err := model.UpdateSystem(h.db, id, u)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, system)
}

// DeleteSystem deletes a specified system in the system.
func (h *Handler) DeleteSystem(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	err := model.DeleteSystem(h.db, id)
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
