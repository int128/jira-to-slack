package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router returns a router for this application.
func Router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &IndexHandler{}).Methods("GET")
	r.Handle("/", handlers.ContentTypeHandler(&WebhookHandler{}, "application/json")).Methods("POST")
	return r
}

// New returns a standalone HTTP server.
func New(addr string) *http.Server {
	m := http.NewServeMux()
	m.Handle("/", handlers.LoggingHandler(os.Stdout, Router()))
	m.Handle("/healthz", &HealthzHandler{})
	return &http.Server{
		Handler: m,
		Addr:    addr,
	}
}
