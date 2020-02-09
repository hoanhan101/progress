package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// System represents a row in system database table.
type System struct {
	ID          string    `db:"system_id" json:"system_id"`
	GoalID      string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	Repeat      string    `db:"repeat" json:"repeat"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// CreateSystem creates a system in the database.
func CreateSystem(db *sqlx.DB, goal_id string, name string, repeat string) (*System, error) {
	s := System{
		ID:          uuid.New().String(),
		GoalID:      goal_id,
		Name:        name,
		Repeat:      repeat,
		DateCreated: time.Now().UTC(),
	}

	const q = `
		INSERT INTO systems
		(system_id, goal_id, name, repeat, date_created)
		VALUES ($1, $2, $3, $4, $5)`

	if _, err := db.Exec(q, s.ID, s.GoalID, s.Name, s.Repeat, s.DateCreated); err != nil {
		return nil, err
	}

	return &s, nil
}

// GetSystems retrieves all systems from the database.
func GetSystems(db *sqlx.DB) ([]System, error) {
	systems := []System{}
	const q = `SELECT * FROM systems;`

	if err := db.Select(&systems, q); err != nil {
		return nil, err
	}

	return systems, nil
}

// GetSystem retrieves a system from the database.
func GetSystem(db *sqlx.DB, id string) (*System, error) {
	var s System

	const q = `
		SELECT *
		FROM systems as s
		WHERE s.system_id = $1`

	if err := db.Get(&s, q, id); err != nil {
		return nil, err
	}

	return &s, nil
}

// UpdateSystem  updates a system from the database.
func UpdateSystem(db *sqlx.DB, id string, name string, repeat string) (*System, error) {
	s, err := GetSystem(db, id)
	if err != nil {
		return nil, err
	}

	// Only update if the given input is a non-zero value. Otherwise, use the
	// existing value.
	if name != "" {
		s.Name = name
	}

	if repeat != "" {
		s.Repeat = repeat
	}

	const q = `
		UPDATE systems SET
		"name" = $2,
		"repeat" = $3
		WHERE system_id = $1`

	if _, err = db.Exec(q, id, s.Name, s.Repeat); err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteSystem deletes a system from the database.
func DeleteSystem(db *sqlx.DB, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	const q = `DELETE FROM systems WHERE system_id = $1`

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}
