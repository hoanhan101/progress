package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/handler"
)

// Start runs the server.
func Start(cfg *config.Config, db *sqlx.DB) error {
	server := echo.New()
	handle := handler.NewHandler(cfg, db)

	// Register middleware.
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Register routes.
	server.GET("/config", handle.Config)
	server.GET("/health", handle.Health)

	server.POST("/goal", handle.CreateGoal)
	server.GET("/goal", handle.GetGoals)
	server.GET("/goal/:id", handle.GetGoal)
	// server.PUT("/goal/:id",handle.UpdateGoal)
	// server.DELETE("/goal/:id",handle.DeleteGoal)

	server.GET("/systems", handle.ListSystems)

	// Start the server.
	err := server.Start(cfg.Server.Address)
	if err != nil {
		server.Logger.Fatal(err)
		return err
	}

	return nil
}
