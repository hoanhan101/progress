package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// System represents a row in system table.
type System struct {
	ID          string    `db:"system_id" json:"system_id"`
	GoalID      string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	Repeat      string    `db:"repeat" json:"repeat"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

// GetSystems retrieves all systems from the table.
func GetSystems(db *sqlx.DB) ([]System, error) {
	systems := []System{}
	const q = `SELECT * FROM systems;`

	if err := db.Select(&systems, q); err != nil {
		return nil, err
	}

	return systems, nil
}
