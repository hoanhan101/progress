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
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// CreateGoal creates a goal in the database.
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

	if _, err := db.Exec(q, g.ID, g.Name, g.DateCreated); err != nil {
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
func UpdateGoal(db *sqlx.DB, id string, name string) (*Goal, error) {
	g, err := GetGoal(db, id)
	if err != nil {
		return nil, err
	}

	g.Name = name

	const q = `
		UPDATE goals SET
		"name" = $2
		WHERE goal_id = $1`

	if _, err = db.Exec(q, id, g.Name); err != nil {
		return nil, err
	}

	return g, nil
}

// DeleteGoal deletes a goal from the database.
func DeleteGoal(db *sqlx.DB, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	const q = `DELETE FROM goals WHERE goal_id = $1`

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}
