package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/int128/jira-to-slack/formatter"
	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/slack"
	"github.com/int128/slack/dialect"
)

// Webhook handles requests from JIRA webhook.
type Webhook struct {
	HTTPClient *http.Client // Default to http.DefaultClient
}

type webhookParams struct {
	webhook  string
	username string
	icon     string
	dialect  dialect.Dialect
	debug    bool
}

func parseWebhookParams(r *http.Request) (*webhookParams, error) {
	var p webhookParams
	q := r.URL.Query()
	p.webhook = q.Get("webhook")
	if p.webhook == "" {
		return nil, fmt.Errorf("missing query parameter. Request with ?webhook=https://hooks.slack.com/xxx")
	}
	p.username = q.Get("username")
	p.icon = q.Get("icon")
	switch q.Get("dialect") {
	case "":
		p.dialect = &dialect.Slack{}
	case "slack":
		p.dialect = &dialect.Slack{}
	case "mattermost":
		p.dialect = &dialect.Mattermost{}
	default:
		return nil, fmt.Errorf("dialect must be slack (default) or mattermost")
	}
	switch q.Get("debug") {
	case "": // default
	case "0": // default
	case "1":
		p.debug = true
	default:
		return nil, fmt.Errorf("debug must be 0 (default) or 1")
	}
	return &p, nil
}

func (h *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	p, err := parseWebhookParams(r)
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
	sc := slack.Client{
		WebhookURL: p.webhook,
		HTTPClient: h.HTTPClient,
	}
	if err := sc.Send(m); err != nil {
		e := fmt.Sprintf("Could not send the message to Slack: %s", err)
		log.Print(e)
		http.Error(w, e, http.StatusInternalServerError)
	}
}
