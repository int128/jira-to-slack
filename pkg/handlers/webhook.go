package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/jira-to-slack/pkg/usecases"
	"github.com/int128/slack/dialect"
	"golang.org/x/xerrors"
)

// Webhook handles requests from JIRA webhook.
type Webhook struct {
	HTTPClientFactory func(context.Context) *http.Client // Default to http.DefaultClient
}

type WebhookParams struct {
	Webhook  string
	Username string
	Channel  string
	Icon     string
	Dialect  dialect.Dialect
	Debug    bool
}

func parseWebhookParams(q url.Values) (*WebhookParams, error) {
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

func (h *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, err := h.Serve(r.Context(), r.URL.Query(), r.Body)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	w.WriteHeader(code)
}

func (h *Webhook) Serve(ctx context.Context, v url.Values, body io.Reader) (int, error) {
	params, err := parseWebhookParams(v)
	if err != nil {
		return http.StatusBadRequest, xerrors.Errorf("could not parse query: %w", err)
	}
	var event jira.Event
	if err := json.NewDecoder(body).Decode(&event); err != nil {
		return http.StatusBadRequest, xerrors.Errorf("could not parse body: %w", err)
	}
	if params.Debug {
		log.Printf("Received parameters %+v", params)
		log.Printf("Received event %+v", event)
	}
	var hc *http.Client
	if h.HTTPClientFactory != nil {
		hc = h.HTTPClientFactory(ctx)
	}

	in := usecases.WebhookIn{
		JiraEvent:       &event,
		SlackWebhookURL: params.Webhook,
		SlackUsername:   params.Username,
		SlackChannel:    params.Channel,
		SlackIcon:       params.Icon,
		SlackDialect:    params.Dialect,
		HTTPClient:      hc,
	}
	var u usecases.Webhook
	if err := u.Do(ctx, in); err != nil {
		return http.StatusInternalServerError, xerrors.Errorf("error: %w", err)
	}
	return http.StatusOK, nil
}
