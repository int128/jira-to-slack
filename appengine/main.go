package main

import (
	"net/http"

	"github.com/int128/jira-to-slack/router"
	"google.golang.org/appengine"
)

func main() {
	http.Handle("/", router.New())
	appengine.Main()
}
