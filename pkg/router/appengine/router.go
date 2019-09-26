package appengine

import (
	"context"
	"net/http"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/int128/jira-to-slack/pkg/handlers"
	"google.golang.org/appengine/urlfetch"
)

func httpClientFactory(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}

func New() http.Handler {
	m := mux.NewRouter()
	m.Handle("/", &handlers.Index{}).Methods("GET")
	m.Handle("/",
		gh.ContentTypeHandler(
			&handlers.Webhook{HTTPClientFactory: httpClientFactory},
			"application/json",
		),
	).Methods("POST")
	return m
}
