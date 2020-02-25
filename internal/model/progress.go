package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Progress represents a row in progress database table.
type Progress struct {
	ID             string    `db:"progress_id" json:"progress_id"`
	SystemID       string    `db:"system_id" json:"system_id"`
	Context        string    `db:"context" json:"context"`
	Completed      bool      `db:"completed" json:"completed"`
	MeasurableData int       `db:"measurable_data" json:"measurable_data"`
	MeasurableUnit string    `db:"measurable_unit" json:"measurable_unit"`
	Sets           int       `db:"sets" json:"sets"`
	Reps           int       `db:"reps" json:"reps"`
	Link           string    `db:"link" json:"link"`
	DateCreated    time.Time `db:"date_created" json:"date_created"`
}

// NewProgress is what required to create a new progress.
type NewProgress struct {
	SystemID       string    `json:"system_id" validate:"required"`
	Context        string    `json:"context"`
	Completed      bool      `json:"completed" validate:"required"`
	MeasurableData int       `json:"measurable_data"`
	MeasurableUnit string    `json:"measurable_unit"`
	Sets           int       `json:"sets"`
	Reps           int       `json:"reps"`
	Link           string    `json:"link"`
	DateCreated    time.Time `json:"date_created"`
}

// CreateProgress creates a progress in the database.
func CreateProgress(db *sqlx.DB, n *NewProgress) (*Progress, error) {
	p := Progress{
		ID:             uuid.New().String(),
		SystemID:       n.SystemID,
		Context:        n.Context,
		Completed:      n.Completed,
		MeasurableData: n.MeasurableData,
		MeasurableUnit: n.MeasurableUnit,
		Link:           n.Link,
	}

	// NOTE - there might be a better way to set defaults.
	if n.DateCreated.IsZero() {
		p.DateCreated = time.Now()
	} else {
		p.DateCreated = n.DateCreated
	}

	if n.Sets <= 0 {
		p.Sets = 1
	} else {
		p.Sets = n.Sets
	}

	if n.Reps <= 0 {
		p.Reps = 1
	} else {
		p.Reps = n.Reps
	}

	const q = `
		INSERT INTO progress
		(progress_id, system_id, context, completed, measurable_data, measurable_unit, sets, reps, link, date_created)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	if _, err := db.Exec(q, p.ID, p.SystemID, p.Context, p.Completed, p.MeasurableData, p.MeasurableUnit, p.Sets, p.Reps, p.Link, p.DateCreated); err != nil {
		return nil, err
	}

	return &p, nil
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
