package formatter

import (
	"bufio"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/int128/jira-to-slack/message"

	"github.com/int128/jira-to-slack/jira"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestFormatJIRAEventToSlackMessage(t *testing.T) {
	matrix := []struct {
		source   string
		expected *message.Message
	}{
		{
			source: "../testdata/issue_created.json",
			expected: &message.Message{
				Text: "<@alice> created the issue: ",
				Attachments: message.Attachments{{
					Title:     "TEST-4: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-4",
					Text:      loremIpsum,
					Timestamp: 1519993052,
				}},
			},
		},
		{
			source: "../testdata/issue_deleted.json",
			expected: &message.Message{
				Text: "<@alice> deleted the issue (assigned to <@alice>): ",
				Attachments: message.Attachments{{
					Title:     "TEST-4: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-4",
					Timestamp: 1519993669,
				}},
			},
		},
		{
			source: "../testdata/issue_updated_assigned.json",
			expected: &message.Message{
				Text: "<@alice> assigned the issue (assigned to <@alice>): ",
				Attachments: message.Attachments{{
					Title:     "TEST-4: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-4",
					Text:      loremIpsum,
					Timestamp: 1519993563,
				}},
			},
		},
		{
			source:   "../testdata/issue_updated_comment_deleted.json",
			expected: nil,
		},
		{
			source: "../testdata/issue_updated_commented.json",
			expected: &message.Message{
				Text: "<@alice> commented the issue: <@bob>",
				Attachments: message.Attachments{{
					Title:     "TEST-4: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-4",
					Text:      loremIpsum,
					Timestamp: 1519993498,
				}},
			},
		},
		{
			source: "../testdata/issue_updated_summary.json",
			expected: &message.Message{
				Text: "<@alice> updated the issue: ",
				Attachments: message.Attachments{{
					Title:     "TEST-2: Lorem Ipsum",
					TitleLink: "https://jira.example.com/browse/TEST-2",
					Timestamp: 1520002692,
				}},
			},
		},
		{
			source:   "../testdata/issue_updated_transition.json",
			expected: nil,
		},
		{
			source:   "../testdata/comment_created.json",
			expected: nil,
		},
		{
			source:   "../testdata/comment_deleted.json",
			expected: nil,
		},
	}
	formatter := &Formatter{&message.SlackDialect{}}
	for i := 0; i < len(matrix); i++ {
		m := matrix[i]
		t.Run(m.source, func(t *testing.T) {
			f, err := os.Open(m.source)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			r := bufio.NewReader(f)
			var event jira.Event
			if err := json.NewDecoder(r).Decode(&event); err != nil {
				t.Fatal(err)
			}

			actual := formatter.JIRAEventToSlackMessage(&event)

			switch {
			case m.expected == nil && actual == nil:
				// OK
			case m.expected == nil && actual != nil:
				t.Errorf("message wants nil but %+v", actual)
			case m.expected != nil && actual == nil:
				t.Errorf("message wants non-nil but nil")
			case m.expected.Text != actual.Text:
				t.Errorf("message.Text wants %s but %s", m.expected.Text, actual.Text)
			case !reflect.DeepEqual(m.expected.Attachments, actual.Attachments):
				t.Errorf("message.Attachments wants %+v but %+v", m.expected.Attachments, actual.Attachments)
			}
		})
	}
}
