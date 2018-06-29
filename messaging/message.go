package messaging

import (
	"fmt"

	"github.com/int128/jira-to-slack/jira"
	"github.com/int128/jira-to-slack/slack"
)

// JIRAEventToSlackMessage formats a JIRA event to a Slack message.
func JIRAEventToSlackMessage(event *jira.Event, dialect slack.Dialect) *slack.Message {
	switch {
	case event.IsIssueCreated():
		return &slack.Message{
			Text: formatTitle(event, "created", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueCommented():
		return &slack.Message{
			Text: formatTitle(event, "commented to", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Comment.Body,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueAssigned():
		return &slack.Message{
			Text: formatTitle(event, "assigned", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("summary"):
		return &slack.Message{
			Text: formatTitle(event, "updated", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueFieldUpdated("description"):
		return &slack.Message{
			Text: formatTitle(event, "updated", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}
	case event.IsIssueDeleted():
		return &slack.Message{
			Text: formatTitle(event, "deleted", dialect),
			Attachments: slack.Attachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}
	default:
		return nil
	}
}

func formatTitle(event *jira.Event, verb string, dialect slack.Dialect) string {
	return fmt.Sprintf("%s %s %s:",
		dialect.Mention(event.User.Name),
		verb,
		formatIssue(event.Issue, dialect))
}

func formatIssue(issue *jira.Issue, dialect slack.Dialect) string {
	if issue.Fields.Assignee == nil {
		return "the issue"
	}
	return fmt.Sprintf("the issue (assigned to %s)",
		dialect.Mention(issue.Fields.Assignee.Name))
}
