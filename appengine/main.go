package main

import (
	"net/http"

	aeRouter "github.com/int128/jira-to-slack/pkg/router/appengine"
	"google.golang.org/appengine"
)

func main() {
	http.Handle("/", aeRouter.New())
	appengine.Main()
}
