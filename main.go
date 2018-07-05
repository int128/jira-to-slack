package main

import (
	"log"
	"net/http"

	"github.com/int128/jira-to-slack/router"
)

func main() {
	addr := ":3000"
	s := &http.Server{
		Handler: router.NewForStandalone(),
		Addr:    addr,
	}
	log.Printf("Listening on %s", addr)
	log.Fatal(s.ListenAndServe())
}
