package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/database"
	"github.com/hoanhan101/progress/internal/goal"
)

// Server is the HTTP API server component.
type Server struct {
	config *config.Config
	server *echo.Echo
	db     *sqlx.DB
}

// NewServer returns a new Server.
func NewServer(cfg *config.Config, d *sqlx.DB) *Server {
	s := &Server{
		config: cfg,
		server: echo.New(),
		db:     d,
	}

	// Register middleware.
	s.server.Use(middleware.Logger())
	s.server.Use(middleware.Recover())

	// Register routes.
	s.server.GET("/config", s.configHandler)
	s.server.GET("/health", s.healthHandler)
	s.server.GET("/goals", s.goalsHandler)

	return s
}

func (s *Server) Run() error {
	err := s.server.Start(s.config.Server.Address)
	if err != nil {
		s.server.Logger.Fatal(err)
		return err
	}

	return nil
}

func (s *Server) configHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.config)
}

func (s *Server) healthHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "ok",
	}

	err := database.StatusCheck(s.db)
	if err != nil {
		resp["status"] = "db not ready"
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) goalsHandler(c echo.Context) error {
	goals, err := goal.List(s.db)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": goals})
}
