package messaging

import (
	"bufio"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/int128/jira-to-slack/jira"

	"github.com/int128/jira-to-slack/slack"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestFormatJIRAEventToSlackMessage(t *testing.T) {
	fixtures := []struct {
		source   string
		expected slack.Message
	}{
		{
			source: "../testdata/issue_created.json",
			expected: slack.Message{
				Text: "<@alice> created the issue:",
				Attachments: slack.Attachments{{
					Title:     "TEST-4: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-4",
					Text:      loremIpsum,
					Timestamp: 1519993052,
				}},
			},
		},
	}
	for i := 0; i < len(fixtures); i++ {
		fixture := fixtures[i]

		f, err := os.Open(fixture.source)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		var event jira.Event
		if err := json.NewDecoder(r).Decode(&event); err != nil {
			t.Fatal(err)
		}

		expected := fixture.expected
		actual := JIRAEventToSlackMessage(&event, slack.Dialect("slack"))
		if actual.Text != expected.Text {
			t.Errorf("Text = %s, want %s", actual.Text, expected.Text)
		}
		if !reflect.DeepEqual(actual.Attachments, expected.Attachments) {
			t.Errorf("Attachments\n actual = %+v,\n expected = %+v", actual.Attachments, expected.Attachments)
		}
	}
}
