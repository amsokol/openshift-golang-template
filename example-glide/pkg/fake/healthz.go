// Package fake provides sample methods
package fake

// Live emulates application liveness probe
// It returns response if success or error if check failed
func Live() (string, error) {
	return "I'm alive!", nil
}

// Ready emulates application readiness probe
// It returns response if success or error if check failed
func Ready() (string, error) {
	return "I'm ready!", nil
}
