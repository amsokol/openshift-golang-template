package main

import (
	"context"

	"os"
	"os/signal"

	"time"

	"net/http"

	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func hello(c echo.Context) error {
	host, err := os.Hostname()
	if err != nil {
		host = err.Error()
	}
	return c.String(http.StatusOK, fmt.Sprintf("Hello World from server %s! Now is %s",
		host,
		time.Now().String()))
}

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "I'm OK!")
}

func main() {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Routes
	e.GET("/", hello)
	e.GET("/healthz", healthz)

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
