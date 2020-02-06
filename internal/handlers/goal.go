package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Goal represents a row in goal table.
type Goal struct {
	GoalID      string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// ListGoals gets all Goal from the database.
func (h *Handler) ListGoals(c echo.Context) error {
	goals := []Goal{}
	const q = `SELECT * FROM goals;`

	if err := h.db.Select(&goals, q); err != nil {
		return errors.Wrap(err, "selecting goals")
	}

	return c.JSON(http.StatusOK, goals)
}
