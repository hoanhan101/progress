package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type System struct {
	SystemID    string    `db:"system_id" json:"system_id"`
	GoalID      string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	Repeat      string    `db:"repeat" json:"repeat"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// ListSystems gets all System from the database.
func (h *Handler) ListSystems(c echo.Context) error {
	systems := []System{}
	const q = `SELECT * FROM systems;`

	if err := h.db.Select(&systems, q); err != nil {
		return errors.Wrap(err, "selecting systems")
	}

	return c.JSON(http.StatusOK, systems)
}
