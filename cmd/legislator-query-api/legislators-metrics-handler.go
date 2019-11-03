package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics - provide a metrics endpoint.
// @TODO: add http_request_total. https://github.com/brancz/prometheus-example-app/blob/master/main.go
// https://github.com/heptiolabs/healthcheck/blob/master/example_test.go
func Metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
