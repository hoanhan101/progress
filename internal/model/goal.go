package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Goal represents a row in goal table.
type Goal struct {
	ID          string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// CreateGoal creates a goal in the table for a given name.
func CreateGoal(db *sqlx.DB, name string) (*Goal, error) {
	g := Goal{
		ID:          uuid.New().String(),
		Name:        name,
		DateCreated: time.Now().UTC(),
	}
	const q = `
		INSERT INTO goals
		(goal_id, name, date_created)
		VALUES ($1, $2, $3)`

	_, err := db.Exec(q, g.ID, g.Name, g.DateCreated)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

// ListGoals gets all goals in the table.
func ListGoals(db *sqlx.DB) ([]Goal, error) {
	goals := []Goal{}
	const q = `SELECT * FROM goals;`

	if err := db.Select(&goals, q); err != nil {
		return nil, err
	}

	return goals, nil
}
