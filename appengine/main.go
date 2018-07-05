package main

import (
	"net/http"

	"github.com/int128/jira-to-slack/server"
	"google.golang.org/appengine"
)

func main() {
	http.Handle("/", server.Router())
	appengine.Main()
}
