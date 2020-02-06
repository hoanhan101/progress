package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// System represents a row in system table.
type System struct {
	ID          string    `db:"system_id" json:"system_id"`
	GoalID      string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	Repeat      string    `db:"repeat" json:"repeat"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// ListSystems list all systems.
func (h *Handler) ListSystems(c echo.Context) error {
	systems := []System{}
	const q = `SELECT * FROM systems;`

	if err := h.db.Select(&systems, q); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to select: %v", err),
		)
	}

	return c.JSON(http.StatusOK, systems)
}
