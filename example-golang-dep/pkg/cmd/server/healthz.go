package server

import (
	"net/http"

	"github.com/amsokol/openshift-golang-template/example-golang-dep/pkg/fake"
)

// live provides HTTP endpoint for application liveness probe
func live(w http.ResponseWriter, r *http.Request) {
	s, err := fake.Live()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.Write([]byte(s))
	}
}

// ready provides HTTP endpoint for application readiness probe
func ready(w http.ResponseWriter, r *http.Request) {
	s, err := fake.Ready()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.Write([]byte(s))
	}
}
