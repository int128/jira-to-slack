package server

import (
	"fmt"
	"net/http"
)

// IndexHandler handles index requests.
type IndexHandler struct{}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
