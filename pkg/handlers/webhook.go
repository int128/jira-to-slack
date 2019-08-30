package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/jira-to-slack/pkg/usecases"
	"github.com/int128/slack/dialect"
)

// Webhook handles requests from JIRA webhook.
type Webhook struct {
	HTTPClientFactory func(*http.Request) *http.Client // Default to http.DefaultClient
}

type WebhookParams struct {
	Webhook  string
	Username string
	Channel  string
	Icon     string
	Dialect  dialect.Dialect
	Debug    bool
}

func ParseWebhookParams(q url.Values) (*WebhookParams, error) {
	var p WebhookParams
	p.Webhook = q.Get("webhook")
	if p.Webhook == "" {
		return nil, fmt.Errorf("missing query parameter. Request with ?webhook=https://hooks.slack.com/xxx")
	}
	p.Username = q.Get("username")
	p.Channel = q.Get("channel")
	p.Icon = q.Get("icon")
	switch q.Get("dialect") {
	case "":
		p.Dialect = &dialect.Slack{}
	case "slack":
		p.Dialect = &dialect.Slack{}
	case "mattermost":
		p.Dialect = &dialect.Mattermost{}
	default:
		return nil, fmt.Errorf("dialect must be slack (default) or mattermost")
	}
	switch q.Get("debug") {
	case "": // default
	case "0": // default
	case "1":
		p.Debug = true
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
	params, err := ParseWebhookParams(r.URL.Query())
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
	if params.Debug {
		log.Printf("Received parameters %+v", params)
		log.Printf("Received event %+v", &event)
	}
	var hc *http.Client
	if h.HTTPClientFactory != nil {
		hc = h.HTTPClientFactory(r)
	}

	in := usecases.WebhookIn{
		JiraEvent:       event,
		SlackWebhookURL: params.Webhook,
		SlackUsername:   params.Username,
		SlackChannel:    params.Channel,
		SlackIcon:       params.Icon,
		SlackDialect:    params.Dialect,
		HTTPClient:      hc,
	}
	var u usecases.Webhook
	if err := u.Do(ctx, in); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
