package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// IndexHandler shows the site index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// HealthzHandler responds to a health check request
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// JIRAWebhookHandler processes a JIRA webhook event
type JIRAWebhookHandler struct {
	SlackWebhookURL string
	Dialect         SlackAPIDialect
	Username        string
	IconEmoji       string
	IconURL         string
}

func (s *JIRAWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var jiraEvent JIRAEvent
	if err := json.NewDecoder(r.Body).Decode(&jiraEvent); err != nil {
		log.Printf("Could not decode the request body: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := FormatJIRAEventToSlackMessage(&jiraEvent, s.Dialect)
	if message == nil {
		message.Username = s.Username
		message.IconEmoji = s.IconEmoji
		message.IconURL = s.IconURL
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := SendSlackMessage(s.SlackWebhookURL, message); err != nil {
		panic(err)
	}
}
