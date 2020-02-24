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
