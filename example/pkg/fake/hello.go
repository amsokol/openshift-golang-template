// Package fake provides sample methods
package fake

import (
	"fmt"
	"os"
	"time"
)

// Hello is simple method
// It returns greetings contains server hostname and date/time
func Hello() (string) {
	host, err := os.Hostname()
	if err != nil {
		host = err.Error()
	}
	return fmt.Sprintf("Hello World from server %s! Now is %s",
		host,
		time.Now().String())
}
