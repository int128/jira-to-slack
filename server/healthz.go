package server

import (
	"fmt"
	"net/http"
)

// HealthzHandler handles health check requests.
type HealthzHandler struct{}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
