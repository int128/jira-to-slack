package main

import (
	"fmt"
	"regexp"
)

// JIRAEvent is a JIRA event sent from a JIRA webhook.
type JIRAEvent struct {
	WebhookEvent string         `json:"webhookEvent"`
	Timestamp    int64          `json:"timestamp"`
	User         *JIRAUser      `json:"user"`
	Issue        *JIRAIssue     `json:"issue"`
	Comment      *JIRAComment   `json:"comment"`
	Changelog    *JIRAChangelog `json:"changelog"`
}

// GetUnixTime returns UNIX time of the event
func (s *JIRAEvent) GetUnixTime() int64 {
	return s.Timestamp / 1000
}

// JIRAUser is a user
type JIRAUser struct {
	Name string `json:"name"`
}

// JIRAIssue is an issue
type JIRAIssue struct {
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields *struct {
		Summary     string    `json:"summary"`
		Description string    `json:"description"`
		Assignee    *JIRAUser `json:"assignee"`
	} `json:"fields"`
}

// FormatKeyAndSummary returns a string like "ISSUE-1: Summary"
func (s *JIRAIssue) FormatKeyAndSummary() string {
	return fmt.Sprintf("%s: %s", s.Key, s.Fields.Summary)
}

// GetURL returns URL of the JIRA issue
func (s *JIRAIssue) GetURL() string {
	base := regexp.MustCompile("/rest/api/.+").ReplaceAllString(s.Self, "")
	return fmt.Sprintf("%s/browse/%s", base, s.Key)
}

// JIRAComment is a comment of an issue
type JIRAComment struct {
	Body string `json:"body"`
}

// JIRAChangelog is a change log of an issue
type JIRAChangelog struct {
	Items []JIRAChangelogItem `json:"items"`
}

// ContainsField returns true if JIRAChangelog has the field of candidates
func (s *JIRAChangelog) ContainsField(candidates ...string) bool {
	for i := 0; i < len(s.Items); i++ {
		for j := 0; j < len(candidates); j++ {
			if s.Items[i].Field == candidates[j] {
				return true
			}
		}
	}
	return false
}

// JIRAChangelogItem is an item of JIRAChangelog
type JIRAChangelogItem struct {
	Field string `json:"field"`
	From  string `json:"fromString"`
	To    string `json:"toString"`
}

const (
	// JIRAEventIssueCreated is sent when an issue is created
	JIRAEventIssueCreated = "jira:issue_created"
	// JIRAEventIssueUpdated is sent when an issue is updated
	JIRAEventIssueUpdated = "jira:issue_updated"
	// JIRAEventIssueDeleted is sent when an issue is deleted
	JIRAEventIssueDeleted = "jira:issue_deleted"
)
