package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hoanhan101/progress/internal/database"
)

// Health is the handler for `/health` endpoint.
func (h *Handler) Health(c echo.Context) error {
	resp := map[string]string{
		"status": "ok",
	}

	err := database.StatusCheck(h.db)
	if err != nil {
		resp["status"] = "db not ready"
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to check database status: %v", err),
		)
	}

	return c.JSON(http.StatusOK, resp)
}
