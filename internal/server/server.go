package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/handlers"
)

// Start runs the server.
func Start(cfg *config.Config, db *sqlx.DB) error {
	server := echo.New()
	handler := handlers.NewHandler(cfg, db)

	// Register middleware.
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Register routes.
	server.GET("/config", handler.Config)
	server.GET("/health", handler.Health)
	server.GET("/goals", handler.ListGoals)
	server.GET("/systems", handler.ListSystems)

	// Start the server.
	err := server.Start(cfg.Server.Address)
	if err != nil {
		server.Logger.Fatal(err)
		return err
	}

	return nil
}
