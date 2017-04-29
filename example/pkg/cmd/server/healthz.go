package server

import (
	"net/http"

	"github.com/amsokol/openshift-golang-template/example/pkg/fake"
)

// healthz provides HTTP endpoint for application healthz check
func healthz(w http.ResponseWriter, r *http.Request) {
	s, err := fake.Healthz()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.Write([]byte(s))
	}
}
