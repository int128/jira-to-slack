package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Message is a message sent to an incoming webhook
type Message struct {
	Username    string      `json:"username,omitempty"`
	IconEmoji   string      `json:"icon_emoji,omitempty"`
	IconURL     string      `json:"icon_url,omitempty"`
	Text        string      `json:"text,omitempty"`
	Attachments Attachments `json:"attachments,omitempty"`
}

// Attachment is an attachment of a message
type Attachment struct {
	Text      string           `json:"text,omitempty"`
	Title     string           `json:"title,omitempty"`
	TitleLink string           `json:"title_link,omitempty"`
	Timestamp int64            `json:"ts,omitempty"`
	Fields    AttachmentFields `json:"fields,omitempty"`
}

// Attachments is an array of Attachment
type Attachments []Attachment

// AttachmentField is a field in the attachment
type AttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// AttachmentFields is an array of AttachmentField
type AttachmentFields []AttachmentField

// Send sends the message to the incoming webhook
func Send(webhookURL string, message *Message) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)
	resp, err := http.Post(webhookURL, "application/json", b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("API error: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}
