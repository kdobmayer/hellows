package main

import "net/http"

type healthResponse struct {
	Status string `json:"status"`
}

// healthHandler provides a basic health check.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{"alive"})
}
