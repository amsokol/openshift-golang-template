package server

import (
	"net/http"

	"github.com/amsokol/openshift-golang-template/example-glide/pkg/fake"
)

// hello provides HTTP endpoint for "Hello" method
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fake.Hello()))
}
