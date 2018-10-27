package main

import (
	"net/http"

	"github.com/int128/jira-to-slack/router"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	hc := urlfetch.Client(ctx)
	router.Core(hc).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	appengine.Main()
}
