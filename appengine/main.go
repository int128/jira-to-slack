package main

import (
	"net/http"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/int128/jira-to-slack/handlers"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func router() http.Handler {
	webhookHandler := &handlers.Webhook{}

	m := mux.NewRouter()
	m.Handle("/", &handlers.Index{}).Methods("GET")
	m.Handle("/", gh.ContentTypeHandler(webhookHandler, "application/json")).Methods("POST")

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		client := urlfetch.Client(ctx)
		webhookHandler.HTTPClient = client
		m.ServeHTTP(w, req)
	})
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
