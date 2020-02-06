package handlers

import (
	"github.com/jmoiron/sqlx"

	"github.com/hoanhan101/progress/internal/config"
)

type Handler struct {
	config *config.Config
	db     *sqlx.DB
}

func NewHandler(cfg *config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		config: cfg,
		db:     db,
	}
}
