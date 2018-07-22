package handlers

import (
	"net/http"
)

const version = "0.0.1"

type versionResponse struct {
	Version string `json:"version"`
}

// Version provides the current version.
func Version(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, versionResponse{version})
}
