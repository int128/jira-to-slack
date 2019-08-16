package main

import (
	"net/http"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/int128/jira-to-slack/pkg/handlers"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func router() http.Handler {
	m := mux.NewRouter()
	m.Handle("/", &handlers.Index{}).Methods("GET")
	m.Handle("/", gh.ContentTypeHandler(&handlers.Webhook{
		HTTPClientFactory: func(r *http.Request) *http.Client {
			return urlfetch.Client(r.Context())
		},
	}, "application/json")).Methods("POST")
	return m
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
