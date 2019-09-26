package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
)

// Index handles index requests.
type Index struct{}

func (h *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, body, err := h.Serve(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	w.WriteHeader(code)
	if _, err := fmt.Fprint(w, body); err != nil {
		log.Printf("could not write body: %s", err)
	}
}

func (h *Index) Serve(v url.Values) (int, string, error) {
	params, err := parseWebhookParams(v)
	if err != nil {
		return http.StatusBadRequest, "", xerrors.Errorf("could not parse query: %w", err)
	}
	return http.StatusOK, fmt.Sprintf("%+v", *params), nil
}
