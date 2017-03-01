// Package main is command line HTTP server
package main

import (
	"net/http"

	"github.com/amsokol/openshift-golang-template/example/pkg/fake"
	"github.com/labstack/echo"
)

// hello provides HTTP endpoint for "Hello" method
func hello(c echo.Context) error {
	return c.String(http.StatusOK, fake.Hello())
}

// healthz provides HTTP endpoint for application healthz check
func healthz(c echo.Context) error {
	s, err := fake.Healthz()
	if err != nil {
		c.Logger().Debugf("Error from health check: %s", err.Error())
		return c.String(http.StatusServiceUnavailable, err.Error())
	}

	c.Logger().Debugf("Success from health check: %s", s)
	return c.String(http.StatusOK, s)
}
