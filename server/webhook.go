package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/jira-to-slack/messaging"
	"github.com/int128/jira-to-slack/slack"
)

// WebhookHandler handles requests from JIRA wehbook.
type WebhookHandler struct{}

type requestParams struct {
	webhook  string
	username string
	icon     string
	dialect  string
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	q := r.URL.Query()
	params := &requestParams{
		q.Get("webhook"),
		q.Get("username"),
		q.Get("icon"),
		q.Get("dialect"),
	}
	if params.webhook == "" {
		e := "Missing query parameter. Request with ?webhook=https://hooks.slack.com/..."
		log.Print(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	var event jira.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		e := fmt.Sprintf("Could not decode the request body: %s", err)
		log.Print(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	m := messaging.JIRAEventToSlackMessage(&event, slack.Dialect("slack"))
	if m == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	m.Username = params.username
	if strings.HasPrefix(params.icon, "http://") || strings.HasPrefix(params.icon, "https://") {
		m.IconURL = params.icon
	} else {
		m.IconEmoji = params.icon
	}
	if err := slack.Send(params.webhook, m); err != nil {
		e := fmt.Sprintf("Could not send the message to Slack: %s", err)
		log.Print(e)
		http.Error(w, e, http.StatusInternalServerError)
	}
}
