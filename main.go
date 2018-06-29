package main

import (
	"log"

	"github.com/int128/jira-to-slack/server"
)

func main() {
	addr := ":3000"
	s := server.New(addr)
	log.Printf("Listening on %s", addr)
	log.Fatal(s.ListenAndServe())
}
