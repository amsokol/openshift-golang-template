package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/", hello)

	// Start server
	port := "8080"
	std := standard.New(":" + port)
	std.SetHandler(e)

	log.Printf("Starting app on port %+v\n", port)
	graceful.ListenAndServe(std.Server, 5*time.Second)
}
