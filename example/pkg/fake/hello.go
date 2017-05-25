// Package fake provides sample methods
package fake

import (
	"os"
	"runtime"
	"strings"
	"time"
)

// Hello is simple method
// It returns greetings contains server hostname and date/time
func Hello() string {
	host, err := os.Hostname()
	if err != nil {
		host = err.Error()
	}
	s := []string{"Hello World from server ", host, "! ",
		"Now is ", time.Now().String(), "! ",
		"Go version ", runtime.Version(), " ", runtime.GOOS, "-", runtime.GOARCH, "!"}
	return strings.Join(s, "")
}
