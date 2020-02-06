package handlers

import (
	"github.com/jmoiron/sqlx"

	"github.com/hoanhan101/progress/internal/config"
)

// Handler holds the configurations and database connection so that all
// handlers can have access to them in 1 place.
type Handler struct {
	config *config.Config
	db     *sqlx.DB
}

// NewHandler returns a new Handler.
func NewHandler(cfg *config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		config: cfg,
		db:     db,
	}
}
