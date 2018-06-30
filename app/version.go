package main

import (
	"net/http"
)

const version = "v1"

type versionResponse struct {
	Version string `json:"version"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, versionResponse{version})
}
