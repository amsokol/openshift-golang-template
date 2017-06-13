package cli

import (
	"fmt"
	"os"
)

// EchoArgs prints arguments are provided via command line
func EchoArgs() {
	if len(os.Args) > 1 {
		fmt.Println("Command line arguments:", os.Args[1:])
	} else {
		fmt.Println("No command line arguments are provided")
	}
}

// GetPort returns HTTP port to listen from environment variable value or default value 8080
func GetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return port
}
