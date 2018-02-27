package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SlackMessage is a message sent to an incoming webhook
type SlackMessage struct {
	Username    string                  `json:"username,omitempty"`
	IconEmoji   string                  `json:"icon_emoji,omitempty"`
	IconURL     string                  `json:"icon_url,omitempty"`
	Text        string                  `json:"text,omitempty"`
	Attachments SlackMessageAttachments `json:"attachments,omitempty"`
}

// SlackMessageAttachment is an attachment of a message
type SlackMessageAttachment struct {
	Text      string                       `json:"text,omitempty"`
	Title     string                       `json:"title,omitempty"`
	TitleLink string                       `json:"title_link,omitempty"`
	Timestamp int64                        `json:"ts,omitempty"`
	Fields    SlackMessageAttachmentFields `json:"fields,omitempty"`
}

// SlackMessageAttachments is an array of SlackMessageAttachment
type SlackMessageAttachments []SlackMessageAttachment

// SlackMessageAttachmentField is a field in the attachment
type SlackMessageAttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// SlackMessageAttachmentFields is an array of SlackMessageAttachmentField
type SlackMessageAttachmentFields []SlackMessageAttachmentField

// FormatMention returns the user mention considering the dialect
func FormatMention(username string, dialect SlackAPIDialect) string {
	switch dialect {
	case "mattermost":
		return fmt.Sprintf("@%s", username)
	default:
		return fmt.Sprintf("<@%s>", username)
	}
}

// SlackAPIDialect represents dialect, i.e. slack or mattermost
type SlackAPIDialect string

// SendSlackMessage sends the message to the incoming webhook
func SendSlackMessage(webhookURL string, message *SlackMessage) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)
	resp, err := http.Post(webhookURL, "application/json", b)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("API error: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}
