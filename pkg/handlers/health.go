package handlers

import "net/http"

type healthResponse struct {
	Status string `json:"status"`
}

// Health provides a basic health check.
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{"alive"})
}
