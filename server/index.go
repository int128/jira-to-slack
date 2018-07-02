package server

import (
	"fmt"
	"net/http"
)

// IndexHandler handles index requests.
type IndexHandler struct{}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := parseWebhookParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Parameter=%+v", p)
	}
}
