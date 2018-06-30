package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Hello from %s!\n", hostname)
}

func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/version", versionHandler)
}

func main() {
	log.Println("Starting application...")

	mux := http.NewServeMux()
	registerHandlers(mux)

	srv := http.Server{Addr: ":8080", Handler: mux}
	go func() {
		log.Printf("Listening on port %d", 8080)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("server: ListenAndServe: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Shutting down...")
	srv.Shutdown(context.Background())
}
