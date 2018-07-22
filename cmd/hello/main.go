package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/kdobmayer/hellows/pkg/handlers"
)

func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.Hello)
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/version", handlers.Version)
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
