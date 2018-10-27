package main

import (
	"log"
	"net/http"

	"github.com/int128/jira-to-slack/router"
)

func main() {
	addr := ":3000"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router.Standalone()))
}
