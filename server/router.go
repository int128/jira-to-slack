package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New returns a new HTTP server.
func New(addr string) *http.Server {
	r := mux.NewRouter()
	r.Handle("/", &IndexHandler{}).Methods("GET")
	r.Handle("/", handlers.ContentTypeHandler(&WebhookHandler{}, "application/json")).Methods("POST")

	m := http.NewServeMux()
	m.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	m.Handle("/healthz", &HealthzHandler{})

	return &http.Server{
		Handler: m,
		Addr:    addr,
	}
}
