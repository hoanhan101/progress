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

// UpdatedProgress is what required to update a progress.
type UpdatedProgress struct {
	Context        string    `json:"context" validate:"required_without=Completed MeasurableData MeasurableUnit Sets Reps Link DateCreated"`
	Completed      bool      `json:"completed" validate:"required_without=Context MeasurableData MeasurableUnit Sets Reps Link DateCreated"`
	MeasurableData int       `json:"measurable_data" validate:"required_without=Context Completed MeasurableUnit Sets Reps Link DateCreated"`
	MeasurableUnit string    `json:"measurable_unit" validate:"required_without=Context Completed MeasurableData Sets Reps Link DateCreated"`
	Sets           int       `json:"sets" validate:"required_without=Context Completed MeasurableData MeasurableUnit Reps Link DateCreated"`
	Reps           int       `json:"reps" validate:"required_without=Context Completed MeasurableData MeasurableUnit Sets Link DateCreated"`
	Link           string    `json:"link" validate:"required_without=Context Completed MeasurableData MeasurableUnit Sets Reps DateCreated"`
	DateCreated    time.Time `json:"date_created" validate:"required_without=Context Completed MeasurableData MeasurableUnit Sets Reps Link"`
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

// GetProgress retrieves a progress from the database.
func GetProgress(db *sqlx.DB, id string) (*Progress, error) {
	var p Progress

	const q = `
		SELECT *
		FROM progress as p
		WHERE p.progress_id = $1`

	if err := db.Get(&p, q, id); err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdateProgress updates a progress from the database.
func UpdateProgress(db *sqlx.DB, id string, u *UpdatedProgress) (*Progress, error) {
	p, err := GetProgress(db, id)
	if err != nil {
		return nil, err
	}

	// Only update if the given value is not a non-zero value.
	if u.Context != "" {
		p.Context = u.Context
	}

	if u.Completed {
		p.Completed = u.Completed
	}

	if u.MeasurableData != 0 {
		p.MeasurableData = u.MeasurableData
	}

	if u.MeasurableUnit != "" {
		p.MeasurableUnit = u.MeasurableUnit
	}

	if u.Sets != 0 {
		p.Sets = u.Sets
	}

	if u.Reps != 0 {
		p.Reps = u.Reps
	}

	if u.Link != "" {
		p.Link = u.Link
	}

	if !u.DateCreated.IsZero() {
		p.DateCreated = u.DateCreated
	}

	const q = `
		UPDATE progress SET
		"context" = $2,
		"completed" = $3,
		"measurable_data" = $4,
		"measurable_unit" = $5,
		"sets" = $6,
		"reps" = $7,
		"link" = $8,
		"date_created" = $9
		WHERE progress_id = $1`

	if _, err = db.Exec(q, id, p.Context, p.Completed, p.MeasurableData, p.MeasurableUnit, p.Sets, p.Reps, p.Link, p.DateCreated); err != nil {
		return nil, err
	}

	return p, nil
}

// DeleteProgress deletes a progress from the database.
func DeleteProgress(db *sqlx.DB, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	const q = `DELETE FROM progress WHERE progress_id = $1`

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}
