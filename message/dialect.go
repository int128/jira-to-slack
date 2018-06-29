package message

import "fmt"

// Dialect represents dialect, i.e. slack or mattermost
type Dialect interface {
	Mention(string) string
}

// SlackDialect is a dialect for Slack API.
type SlackDialect struct{}

// Mention returns the user mention considering the dialect.
func (s *SlackDialect) Mention(username string) string {
	return fmt.Sprintf("<@%s>", username)
}

// MattermostDialect is a dialect for Mattermost API.
type MattermostDialect struct{}

// Mention returns the user mention considering the dialect.
func (s *MattermostDialect) Mention(username string) string {
	return fmt.Sprintf("@%s", username)
}
