// Package fake provides sample methods
package fake

// Healthz emulates application health check
// It returns response if success or error if check failed
func Healthz() (string, error) {
	return "I'm OK!", nil
}
