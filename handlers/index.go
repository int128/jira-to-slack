package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Index handles index requests.
type Index struct{}

func (h *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := parseWebhookParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := fmt.Fprintf(w, "Parameter=%+v", p); err != nil {
		log.Printf("Error while writing response: %s", err)
	}
}
