package main

import (
	"context"

	"os"
	"os/signal"

	"time"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Routes
	e.GET("/", hello)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
