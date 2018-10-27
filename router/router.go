package router

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Core returns a router with index and webhook.
func Core(hc *http.Client) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &IndexHandler{}).Methods("GET")
	r.Handle("/", handlers.ContentTypeHandler(&WebhookHandler{hc}, "application/json")).Methods("POST")
	return r
}

// Standalone returns a router with logging and health check endpoint.
func Standalone() http.Handler {
	r := http.NewServeMux()
	r.Handle("/", handlers.LoggingHandler(os.Stdout, Core(nil)))
	r.Handle("/healthz", &HealthzHandler{})
	return r
}
