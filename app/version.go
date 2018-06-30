package main

import (
	"net/http"
)

const version = "latest"

type versionResponse struct {
	Version string `json:"version"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, versionResponse{version})
}
