package handlers

import (
	"log"
	"net/http"
)

// Healthz handles health check requests.
type Healthz struct{}

func (h *Healthz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Error while writing response: %s", err)
	}
}
