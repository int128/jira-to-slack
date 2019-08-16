package usecases

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/int128/jira-to-slack/pkg/formatter"
	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/slack"
	"github.com/int128/slack/dialect"
)

type Webhook struct {
}

type WebhookIn struct {
	JiraEvent       *jira.Event
	SlackWebhookURL string
	SlackUsername   string
	SlackIcon       string
	SlackDialect    dialect.Dialect
	HTTPClient      *http.Client // default to http.DefaultClient
}

func (u *Webhook) Do(ctx context.Context, in WebhookIn) error {
	f := formatter.New(in.SlackDialect)
	message := f.JIRAEventToSlackMessage(in.JiraEvent)
	if message == nil {
		return nil
	}

	message.Username = in.SlackUsername
	if strings.HasPrefix(in.SlackIcon, "http://") || strings.HasPrefix(in.SlackIcon, "https://") {
		message.IconURL = in.SlackIcon
	} else {
		message.IconEmoji = in.SlackIcon
	}

	sc := slack.Client{
		WebhookURL: in.SlackWebhookURL,
		HTTPClient: in.HTTPClient,
	}
	if err := sc.Send(message); err != nil {
		return fmt.Errorf("could not send the message to Slack: %s", err)
	}
	return nil
}
