package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/jira-to-slack/pkg/usecases"
	"github.com/int128/slack/dialect"
)

// Webhook handles requests from JIRA webhook.
type Webhook struct {
	HTTPClientFactory func(*http.Request) *http.Client // Default to http.DefaultClient
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

func parseWebhookBody(r *http.Request) (*jira.Event, error) {
	var event jira.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, fmt.Errorf("could not decode json of request body: %s", err)
	}
	return &event, nil
}

func (h *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()
	params, err := parseWebhookParams(r)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := parseWebhookBody(r)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if params.debug {
		log.Printf("Received parameters %+v", params)
		log.Printf("Received event %+v", &event)
	}
	var hc *http.Client
	if h.HTTPClientFactory != nil {
		hc = h.HTTPClientFactory(r)
	}

	in := usecases.WebhookIn{
		JiraEvent:       event,
		SlackWebhookURL: params.webhook,
		SlackUsername:   params.username,
		SlackIcon:       params.icon,
		SlackDialect:    params.dialect,
		HTTPClient:      hc,
	}
	var u usecases.Webhook
	if err := u.Do(ctx, in); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
