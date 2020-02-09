package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/handler"
)

// Start runs the server.
func Start(cfg *config.Config, db *sqlx.DB) error {
	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
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
	server.PUT("/goal/:id", handle.UpdateGoal)
	server.DELETE("/goal/:id", handle.DeleteGoal)

	server.POST("/system", handle.CreateSystem)
	server.GET("/system", handle.GetSystems)
	server.GET("/system/:id", handle.GetSystem)
	server.PUT("/system/:id", handle.UpdateSystem)
	server.DELETE("/system/:id", handle.DeleteSystem)

	server.GET("/progress", handle.GetProgresses)

	// Start the server.
	err := server.Start(cfg.Server.Address)
	if err != nil {
		server.Logger.Fatal(err)
		return err
	}

	return nil
}
