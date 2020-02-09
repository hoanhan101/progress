package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"

	"github.com/hoanhan101/progress/internal/config"
)

// Handler holds the configurations and database connection so that all
// handlers can have access to them in 1 place.
type Handler struct {
	config    *config.Config
	db        *sqlx.DB
	validator *validator.Validate
}

// NewHandler returns a new Handler.
func NewHandler(cfg *config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		config:    cfg,
		db:        db,
		validator: validator.New(),
	}
}
