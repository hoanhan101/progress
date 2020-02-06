package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/database"
)

func (h *Handler) Health(c echo.Context) error {
	resp := map[string]string{
		"status": "ok",
	}

	err := database.StatusCheck(h.db)
	if err != nil {
		resp["status"] = "db not ready"
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
