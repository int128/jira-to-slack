package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/int128/jira-to-slack/formatter"
	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/jira-to-slack/message"
)

// WebhookHandler handles requests from JIRA wehbook.
type WebhookHandler struct{}

type requestParams struct {
	webhook  string
	username string
	icon     string
	dialect  message.Dialect
	debug    bool
}

func parseRequestParams(r *http.Request) (*requestParams, error) {
	p := &requestParams{}
	q := r.URL.Query()
	p.webhook = q.Get("webhook")
	if p.webhook == "" {
		return nil, fmt.Errorf("Missing query parameter. Request with ?webhook=https://hooks.slack.com/xxx")
	}
	p.username = q.Get("username")
	p.icon = q.Get("icon")
	switch q.Get("dialect") {
	case "":
		p.dialect = &message.SlackDialect{}
	case "slack":
		p.dialect = &message.SlackDialect{}
	case "mattermost":
		p.dialect = &message.MattermostDialect{}
	default:
		return nil, fmt.Errorf("dialect must be slack (default) or mattermost")
	}
	switch q.Get("debug") {
	case "1":
		p.debug = true
	}
	return p, nil
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p, err := parseRequestParams(r)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var event jira.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		e := fmt.Sprintf("Could not decode the request body: %s", err)
		log.Print(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	if p.debug {
		log.Printf("Received parameters %+v", p)
		log.Printf("Received event %+v", &event)
	}

	f := formatter.New(p.dialect)
	m := f.JIRAEventToSlackMessage(&event)
	if m == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	m.Username = p.username
	if strings.HasPrefix(p.icon, "http://") || strings.HasPrefix(p.icon, "https://") {
		m.IconURL = p.icon
	} else {
		m.IconEmoji = p.icon
	}
	if p.debug {
		log.Printf("Sending %+v", m)
	}
	if err := message.Send(p.webhook, m); err != nil {
		e := fmt.Sprintf("Could not send the message to Slack: %s", err)
		log.Print(e)
		http.Error(w, e, http.StatusInternalServerError)
	}
}
