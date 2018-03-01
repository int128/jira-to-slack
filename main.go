package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	if os.Getenv("SLACK_WEBHOOK") == "" {
		log.Fatal("Set the Slack Incoming Webhook URL by environment variable SLACK_WEBHOOK")
	}

	jiraWebhookHandler := &JIRAWebhookHandler{
		SlackWebhookURL: os.Getenv("SLACK_WEBHOOK"),
		Dialect:         SlackAPIDialect(os.Getenv("SLACK_API_DIALECT")),
		Username:        os.Getenv("SLACK_USERNAME"),
		IconEmoji:       os.Getenv("SLACK_ICON_EMOJI"),
		IconURL:         os.Getenv("SLACK_ICON_URL"),
	}

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.Handle("/", handlers.ContentTypeHandler(jiraWebhookHandler, "application/json")).Methods("POST")

	m := http.NewServeMux()
	m.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	m.HandleFunc("/healthz", HealthzHandler)

	s := http.Server{
		Handler: m,
		Addr:    ":3000",
	}
	log.Fatal(s.ListenAndServe())
}
