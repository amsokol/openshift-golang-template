package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"syscall"

	"github.com/pressly/chi"
)

// Start starts HTTP server
func Start(port string) error {
	logger := log.New(os.Stdout, "", 0)

	r := chi.NewRouter()
	r.Get("/", hello)
	r.Get("/healthz", healthz)

	h := &http.Server{Addr: ":" + port, Handler: r}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-quit
		logger.Println("\nShutting down the server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := h.Shutdown(ctx); err != nil {
			logger.Println("Error when stopping the server: " + err.Error())
		} else {
			logger.Println("Server gracefully stopped")
		}
	}()

	logger.Printf("Listening on http://0.0.0.0:%s\n", port)
	if err := h.ListenAndServe(); err != nil {
		logger.Println("Listening returns error: " + err.Error())
		return err
	}

	logger.Println("\nServer stopped")
	return nil
}
