package formatter

import (
	"fmt"

	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/jira-to-slack/message"
)

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
			Text: f.title(event, "created"),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueCommented():
		return &message.Message{
			Text: f.title(event, "commented to"),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Comment.Body,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueAssigned():
		return &message.Message{
			Text: f.title(event, "assigned"),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("summary"):
		return &message.Message{
			Text: f.title(event, "updated"),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("description"):
		return &message.Message{
			Text: f.title(event, "updated"),
			Attachments: message.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueDeleted():
		return &message.Message{
			Text: f.title(event, "deleted"),
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

func (f *Formatter) title(event *jira.Event, verb string) string {
	return fmt.Sprintf("%s %s %s:",
		f.Dialect.Mention(event.User.Name),
		verb,
		f.issue(event.Issue))
}

func (f *Formatter) issue(issue *jira.Issue) string {
	switch {
	case issue.Fields.Assignee == nil:
		return "the issue"
	default:
		return fmt.Sprintf("the issue (assigned to %s)", f.Dialect.Mention(issue.Fields.Assignee.Name))
	}
}
