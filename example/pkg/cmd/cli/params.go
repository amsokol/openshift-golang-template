package cli

import (
	"os"
)

// GetPort returns HTTP port to listen from environment varible value or default value 8080
func GetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return port
}
