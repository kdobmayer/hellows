package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

// writeJSON writes the value v to the http response stream as json.
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
