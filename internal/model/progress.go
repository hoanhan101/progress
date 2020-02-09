package model

import (
	"time"

	// "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Progress represents a row in progress database table.
type Progress struct {
	ID             string    `db:"progress_id" json:"progress_id"`
	SystemID       string    `db:"system_id" json:"system_id"`
	DateCreated    time.Time `db:"date_created" json:"date_created"`
	Summary        string    `db:"summary" json:"summary"`
	Completed      bool      `db:"completed" json:"completed"`
	MeasurableData int       `db:"measurable_data" json:"measurable_data"`
	MeasurableUnit string    `db:"measurable_unit" json:"measurable_unit"`
	Sets           int       `db:"sets" json:"sets"`
	Reps           int       `db:"reps" json:"reps"`
	Link           string    `db:"link" json:"link"`
}

// GetProgresses retrieves all progress from the database.
func GetProgresses(db *sqlx.DB) ([]Progress, error) {
	progress := []Progress{}
	const q = `SELECT * FROM progress;`

	if err := db.Select(&progress, q); err != nil {
		return nil, err
	}

	return progress, nil
}
