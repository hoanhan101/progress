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

// NewSystem is what required to create a new system.
type NewSystem struct {
	GoalID string `json:"goal_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Repeat string `json:"repeat"`
}

// UpdatedSystem is what required to update a system.
type UpdatedSystem struct {
	Name   string `json:"name" validate:"required_without=Repeat"`
	Repeat string `json:"repeat" validate:"required_without=Name"`
}

// CreateSystem creates a system in the database.
func CreateSystem(db *sqlx.DB, n *NewSystem) (*System, error) {
	s := System{
		ID:          uuid.New().String(),
		GoalID:      n.GoalID,
		Name:        n.Name,
		Repeat:      n.Repeat,
		DateCreated: time.Now(),
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
func UpdateSystem(db *sqlx.DB, id string, u *UpdatedSystem) (*System, error) {
	s, err := GetSystem(db, id)
	if err != nil {
		return nil, err
	}

	// Only update if the given value is not a non-zero value.
	if u.Name != "" {
		s.Name = u.Name
	}

	if u.Repeat != "" {
		s.Repeat = u.Repeat
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
