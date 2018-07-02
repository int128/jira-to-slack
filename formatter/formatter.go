package formatter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/jira-to-slack/message"
)

var jiraMention = regexp.MustCompile(`\[~(\w+)\]|@(\w+)`)

// Formatter performs JIRA and Slack message conversion.
type Formatter struct {
	Dialect message.Dialect
}

// New returns a new Formatter.
func New(dialect message.Dialect) *Formatter {
	return &Formatter{dialect}
}

// JIRAEventToSlackMessage formats a JIRA event to a Slack message.
func (f *Formatter) JIRAEventToSlackMessage(event *jira.Event) *message.Message {
	switch {
	case event.IsIssueCreated():
		return &message.Message{
			Text: f.title(event, "created", f.mentions(event.Issue.Fields.Description)),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueCommented():
		return &message.Message{
			Text: f.title(event, "commented to", f.mentions(event.Comment.Body)),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Comment.Body,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueAssigned():
		return &message.Message{
			Text: f.title(event, "assigned", f.mentions(event.Issue.Fields.Description)),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("summary"):
		return &message.Message{
			Text: f.title(event, "updated", ""),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("description"):
		return &message.Message{
			Text: f.title(event, "updated", f.mentions(event.Issue.Fields.Description)),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueDeleted():
		return &message.Message{
			Text: f.title(event, "deleted", ""),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}
	default:
		return nil
	}
}

// title returns a message title for the JIRA event.
func (f *Formatter) title(event *jira.Event, verb string, additionalMentions string) string {
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

// mentions returns all mentions in the text.
func (f *Formatter) mentions(text string) string {
	all := jiraMention.FindAllStringSubmatch(text, -1)
	if all == nil {
		return ""
	}
	mentions := make([]string, 0)
	for _, m := range all {
		if len(m) == 1 {
			name := m[1]
			mentions = append(mentions, fmt.Sprintf("%s", f.Dialect.Mention(name)))
		}
	}
	return strings.Join(mentions, ", ")
}
