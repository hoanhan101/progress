package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// CreateProgress create a progress in the system.
func (h *Handler) CreateProgress(c echo.Context) error {
	n := new(model.NewProgress)
	if err := c.Bind(n); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(n); err != nil {
		return errBadRequest(err)
	}

	progress, err := model.CreateProgress(h.db, n)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, progress)
}

// GetProgresses gets all progress in the system.
func (h *Handler) GetProgresses(c echo.Context) error {
	progress, err := model.GetProgresses(h.db)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, progress)
}

// GetProgress gets a specified progress in the system.
func (h *Handler) GetProgress(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	progress, err := model.GetProgress(h.db, id)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, progress)
}

// UpdateProgress updates a specified progress in the system.
func (h *Handler) UpdateProgress(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	u := new(model.UpdatedProgress)
	if err := c.Bind(u); err != nil {
		return errBadRequest(err)
	}

	if err := h.validator.Struct(u); err != nil {
		return errBadRequest(err)
	}

	progress, err := model.UpdateProgress(h.db, id, u)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, progress)
}

// DeleteProgress deletes a specified progress in the system.
func (h *Handler) DeleteProgress(c echo.Context) error {
	id := c.Param("id")
	if err := h.validator.Var(id, "required"); err != nil {
		return errBadRequest(err)
	}

	err := model.DeleteProgress(h.db, id)
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
