package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Goal represents a row in goal table.
type Goal struct {
	ID          string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// CreateGoal create a goal.
func (h *Handler) CreateGoal(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"require name value",
		)
	}

	g := Goal{
		ID:          uuid.New().String(),
		Name:        name,
		DateCreated: time.Now().UTC(),
	}
	const q = `
		INSERT INTO goals
		(goal_id, name, date_created)
		VALUES ($1, $2, $3)`

	_, err := h.db.Exec(q, g.ID, g.Name, g.DateCreated)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to insert: %v", err),
		)
	}

	return c.JSON(http.StatusOK, g)
}

// ListGoals lists all goals.
func (h *Handler) ListGoals(c echo.Context) error {
	goals := []Goal{}
	const q = `SELECT * FROM goals;`

	if err := h.db.Select(&goals, q); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to select: %v", err),
		)
	}

	return c.JSON(http.StatusOK, goals)
}
