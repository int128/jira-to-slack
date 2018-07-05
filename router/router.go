package router

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New returns a router for this application.
func New() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &IndexHandler{}).Methods("GET")
	r.Handle("/", handlers.ContentTypeHandler(&WebhookHandler{}, "application/json")).Methods("POST")
	return r
}

// NewForStandalone returns a router with logging and health check endpoint.
func NewForStandalone() http.Handler {
	r := http.NewServeMux()
	r.Handle("/", handlers.LoggingHandler(os.Stdout, New()))
	r.Handle("/healthz", &HealthzHandler{})
	return r
}
