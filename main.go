package main

import (
	"log"
	"net/http"
	"os"

	"github.com/int128/jira-to-slack/pkg/router"
)

var version string

const defaultPort = "3000"

func main() {
	log.Printf("jira-to-slack %s", version)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	addr := ":" + port

	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, router.New()); err != nil {
		log.Fatalf("Error while listening on %s: %s", addr, err)
	}
}
