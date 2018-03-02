package main

import (
	"fmt"
)

// FormatJIRAEventToSlackMessage formats a JIRA event to a Slack message
func FormatJIRAEventToSlackMessage(event *JIRAEvent, dialect SlackAPIDialect) *SlackMessage {
	switch {
	case event.IsIssueCreated():
		return &SlackMessage{
			Text: formatTitle(event, "created", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}

	case event.IsIssueCommented():
		return &SlackMessage{
			Text: formatTitle(event, "commented to", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Comment.Body,
				Timestamp: event.GetUnixTime(),
			}},
		}

	case event.IsIssueAssigned():
		return &SlackMessage{
			Text: formatTitle(event, "assigned", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}

	case event.IsIssueFieldUpdated("summary"):
		return &SlackMessage{
			Text: formatTitle(event, "updated", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}

	case event.IsIssueFieldUpdated("description"):
		return &SlackMessage{
			Text: formatTitle(event, "updated", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Text:      event.Issue.Fields.Description,
				Timestamp: event.GetUnixTime(),
			}},
		}

	case event.IsIssueDeleted():
		return &SlackMessage{
			Text: formatTitle(event, "deleted", dialect),
			Attachments: SlackMessageAttachments{{
				Title:     event.Issue.FormatKeyAndSummary(),
				TitleLink: event.Issue.GetURL(),
				Timestamp: event.GetUnixTime(),
			}},
		}

	default:
		return nil
	}
}

func formatTitle(event *JIRAEvent, verb string, dialect SlackAPIDialect) string {
	return fmt.Sprintf("%s %s %s:",
		dialect.FormatMention(event.User.Name),
		verb,
		formatIssue(event.Issue, dialect))
}

func formatIssue(issue *JIRAIssue, dialect SlackAPIDialect) string {
	if issue.Fields.Assignee == nil {
		return "the issue"
	}
	return fmt.Sprintf("the issue (assigned to %s)", dialect.FormatMention(issue.Fields.Assignee.Name))
}
