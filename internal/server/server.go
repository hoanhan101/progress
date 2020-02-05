package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hoanhan101/progress/internal/config"
)

// Server is the HTTP API server component.
type Server struct {
	config *config.Config
	server *echo.Echo
}

// NewServer returns a new Server.
func NewServer(cfg *config.Config) *Server {
	s := &Server{
		config: cfg,
		server: echo.New(),
	}

	// Register middleware.
	s.server.Use(middleware.Logger())
	s.server.Use(middleware.Recover())

	// Register routes.
	s.server.GET("/", s.helloHandler)
	s.server.GET("/config", s.configHandler)

	return s
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.config.Server.Address))
}

func (s *Server) helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "Hello, World!"})
}

func (s *Server) configHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.config)
}
