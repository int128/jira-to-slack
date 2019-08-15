package formatter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/slack"
	"github.com/int128/slack/dialect"
)

// Formatter performs JIRA and Slack message conversion.
type Formatter struct {
	Dialect dialect.Dialect
}

// New returns a new Formatter.
func New(d dialect.Dialect) *Formatter {
	return &Formatter{d}
}

// JIRAEventToSlackMessage formats a JIRA event to a Slack message.
func (f *Formatter) JIRAEventToSlackMessage(event *jira.Event) *slack.Message {
	switch {
	case event.IsIssueCreated():
		return &slack.Message{
			Text: f.text(event, "created", f.mentions(event.Issue.Fields.Description)),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.UnixTime(),
			}},
		}
	case event.IsIssueCommented():
		return &slack.Message{
			Text: f.text(event, "commented to", f.mentions(event.Comment.Body)),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Text:      event.Comment.Body,
				Timestamp: event.UnixTime(),
			}},
		}
	case event.IsIssueAssigned():
		return &slack.Message{
			Text: f.text(event, "assigned", f.mentions(event.Issue.Fields.Description)),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.UnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("summary"):
		return &slack.Message{
			Text: f.text(event, "updated", ""),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Timestamp: event.UnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("description"):
		return &slack.Message{
			Text: f.text(event, "updated", f.mentions(event.Issue.Fields.Description)),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.UnixTime(),
			}},
		}
	case event.IsIssueDeleted():
		return &slack.Message{
			Text: f.text(event, "deleted", ""),
			Attachments: []slack.Attachment{{
				Title:     f.title(event),
				TitleLink: event.Issue.BrowserURL(),
				Timestamp: event.UnixTime(),
			}},
		}
	default:
		return nil
	}
}

// title returns an attachment title for the JIRA event.
func (f *Formatter) title(event *jira.Event) string {
	return fmt.Sprintf("%s: %s", event.Issue.Key, event.Issue.Fields.Summary)
}

// text returns a message text for the JIRA event.
func (f *Formatter) text(event *jira.Event, verb string, additionalMentions string) string {
	switch {
	case event.Issue.Fields.Assignee.Name == "":
		return fmt.Sprintf("%s %s the issue: %s",
			f.Dialect.Mention(event.User.Name),
			verb,
			additionalMentions)
	default:
		return fmt.Sprintf("%s %s the issue (assigned to %s): %s",
			f.Dialect.Mention(event.User.Name),
			verb,
			f.Dialect.Mention(event.Issue.Fields.Assignee.Name),
			additionalMentions)
	}
}

var jiraMention = regexp.MustCompile(`\[~(\w+)\]|@(\w+)`)

// mentions returns all mentions in the text.
func (f *Formatter) mentions(text string) string {
	all := jiraMention.FindAllStringSubmatch(text, -1)
	if all == nil {
		return ""
	}
	mentions := make([]string, 0)
	for _, m := range all {
		switch {
		case len(m) == 0:
			// something wrong
		case m[2] != "":
			mentions = append(mentions, f.Dialect.Mention(m[2]))
		case m[1] != "":
			mentions = append(mentions, f.Dialect.Mention(m[1]))
		}
	}
	return strings.Join(mentions, ", ")
}
