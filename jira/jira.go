package jira

import (
	"fmt"
	"regexp"
)

// Event is a JIRA event sent from a JIRA webhook.
type Event struct {
	WebhookEvent   string     `json:"webhookEvent"`
	IssueEventType string     `json:"issue_event_type_name"`
	Timestamp      int64      `json:"timestamp"`
	User           *User      `json:"user"`
	Issue          *Issue     `json:"issue"`
	Comment        *Comment   `json:"comment"`
	Changelog      *Changelog `json:"changelog"`
}

// IsIssueCreated returns true when an issue is created
func (s *Event) IsIssueCreated() bool {
	return s.WebhookEvent == "jira:issue_created"
}

// IsIssueCommented is sent when an comment is created
func (s *Event) IsIssueCommented() bool {
	return s.WebhookEvent == "jira:issue_updated" && s.IssueEventType == "issue_commented"
}

// IsIssueAssigned is sent when the issue is assigned
func (s *Event) IsIssueAssigned() bool {
	return s.WebhookEvent == "jira:issue_updated" && s.IssueEventType == "issue_assigned"
}

// IsIssueFieldUpdated is sent when the issue is updated
func (s *Event) IsIssueFieldUpdated(fields ...string) bool {
	return s.WebhookEvent == "jira:issue_updated" && s.Changelog != nil && s.Changelog.ContainsField(fields...)
}

// IsIssueDeleted is sent when an issue is deleted
func (s *Event) IsIssueDeleted() bool {
	return s.WebhookEvent == "jira:issue_deleted"
}

// GetUnixTime returns UNIX time of the event
func (s *Event) GetUnixTime() int64 {
	return s.Timestamp / 1000
}

// User is a user
type User struct {
	Name string `json:"name"`
}

// Issue is an issue
type Issue struct {
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields *struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Assignee    *User  `json:"assignee"`
	} `json:"fields"`
}

// FormatKeyAndSummary returns a string like "ISSUE-1: Summary"
func (s *Issue) FormatKeyAndSummary() string {
	return fmt.Sprintf("%s: %s", s.Key, s.Fields.Summary)
}

// GetURL returns URL of the JIRA issue
func (s *Issue) GetURL() string {
	base := regexp.MustCompile("/rest/api/.+").ReplaceAllString(s.Self, "")
	return fmt.Sprintf("%s/browse/%s", base, s.Key)
}

// Comment is a comment of an issue
type Comment struct {
	Body string `json:"body"`
}

// Changelog is a change log of an issue
type Changelog struct {
	Items []ChangelogItem `json:"items"`
}

// ContainsField returns true if Changelog has the field of candidates
func (s *Changelog) ContainsField(candidates ...string) bool {
	for i := 0; i < len(s.Items); i++ {
		for j := 0; j < len(candidates); j++ {
			if s.Items[i].Field == candidates[j] {
				return true
			}
		}
	}
	return false
}

// ChangelogItem is an item of Changelog
type ChangelogItem struct {
	Field string `json:"field"`
	From  string `json:"fromString"`
	To    string `json:"toString"`
}
