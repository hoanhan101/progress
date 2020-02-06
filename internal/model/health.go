package model

import (
	"github.com/jmoiron/sqlx"

	"github.com/hoanhan101/progress/internal/database"
)

// Health represents a health information.
type Health struct {
	Status string `json:"status"`
}

// GetHealth gets health information for the database.
func GetHealth(db *sqlx.DB) (*Health, error) {
	h := Health{
		Status: "ok",
	}

	err := database.StatusCheck(db)
	if err != nil {
		h.Status = "db not ready"
		return nil, err
	}

	return &h, nil
}
