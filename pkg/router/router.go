package router

import (
	"net/http"
	"os"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/int128/jira-to-slack/pkg/handlers"
)

func New() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &handlers.Index{}).Methods("GET")
	r.Handle("/", gh.ContentTypeHandler(&handlers.Webhook{}, "application/json")).Methods("POST")

	m := http.NewServeMux()
	m.Handle("/", gh.LoggingHandler(os.Stdout, r))
	m.Handle("/healthz", &handlers.Healthz{})
	return m
}
