package handlers

import (
	"fmt"
	"net/http"
	"os"
)

// Hello greets the client with a message containing the server's hostname.
func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Hello from %s!\n", hostname)
}
