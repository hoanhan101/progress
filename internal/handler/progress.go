package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/model"
)

// GetProgresses gets all progress in the system.
func (h *Handler) GetProgresses(c echo.Context) error {
	progress, err := model.GetProgresses(h.db)
	if err != nil {
		return errInternalServer(err)
	}

	return c.JSON(http.StatusOK, progress)
}
