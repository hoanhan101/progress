package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Goal represents a row in goal database table.
type Goal struct {
	ID          string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	Context     string    `db:"context" json:"context"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// NewGoal is what required to create a new goal.
type NewGoal struct {
	Name    string `json:"name" validate:"required"`
	Context string `json:"context"`
}

// UpdatedGoal is what required to update a goal.
type UpdatedGoal struct {
	Name    string `json:"name" validate:"required_without=Context"`
	Context string `json:"context" validate:"required_without=Name"`
}

// CreateGoal creates a goal in the database.
func CreateGoal(db *sqlx.DB, n *NewGoal) (*Goal, error) {
	g := Goal{
		ID:          uuid.New().String(),
		Name:        n.Name,
		Context:     n.Context,
		DateCreated: time.Now().UTC(),
	}

	const q = `
		INSERT INTO goals
		(goal_id, name, context, date_created)
		VALUES ($1, $2, $3, $4)`

	if _, err := db.Exec(q, g.ID, g.Name, g.Context, g.DateCreated); err != nil {
		return nil, err
	}

	return &g, nil
}

// GetGoals retrieves all goals from the database.
func GetGoals(db *sqlx.DB) ([]Goal, error) {
	gs := []Goal{}
	const q = `SELECT * FROM goals;`

	if err := db.Select(&gs, q); err != nil {
		return nil, err
	}

	return gs, nil
}

// GetGoal retrieves a goal from the database.
func GetGoal(db *sqlx.DB, id string) (*Goal, error) {
	var g Goal

	const q = `
		SELECT * 
		FROM goals as g
		WHERE g.goal_id = $1`

	if err := db.Get(&g, q, id); err != nil {
		return nil, err
	}

	return &g, nil
}

// UpdateGoal updates a goal from the database.
func UpdateGoal(db *sqlx.DB, id string, u *UpdatedGoal) (*Goal, error) {
	g, err := GetGoal(db, id)
	if err != nil {
		return nil, err
	}

	// Only update if the given value is not a non-zero value.
	if u.Name != "" {
		g.Name = u.Name
	}

	if u.Context != "" {
		g.Context = u.Context
	}

	const q = `
		UPDATE goals SET
		"name" = $2,
		"context" = $3
		WHERE goal_id = $1`

	if _, err = db.Exec(q, id, g.Name, g.Context); err != nil {
		return nil, err
	}

	return g, nil
}

// DeleteGoal deletes a goal from the database.
func DeleteGoal(db *sqlx.DB, id string) error {
	_, err := GetGoal(db, id)
	if err != nil {
		return err
	}

	const q = `DELETE FROM goals WHERE goal_id = $1`

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}
