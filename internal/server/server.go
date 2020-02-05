package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Echo() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "Hello, World!"})
}
