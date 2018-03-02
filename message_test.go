package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

var fixtures = []struct {
	eventJSON string
	expected  SlackMessage
}{
	{
		eventJSON: `{
			"id": 2,
			"timestamp": 1519952323652,
			"issue": {
				"expand": "renderedFields,names,schema,transitions,operations,editmeta,changelog",
				"id": "99291",
				"self": "https://jira.atlassian.com/rest/api/2/issue/99291",
				"key": "JRA-20002",
				"fields": {
					"summary": "I feel the need for speed",
					"created": "2009-12-16T23:46:10.612-0600",
					"description": "Make the issue nav load 10x faster",
					"labels": [
						"UI",
						"dialogue",
						"move"
					],
					"priority": "Minor"
				}
			},
			"user": {
				"self": "https://jira.atlassian.com/rest/api/2/user?username=brollins",
				"name": "brollins",
				"emailAddress": "bryansemail at atlassian dot com",
				"avatarUrls": {
					"16x16": "https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
					"48x48": "https://jira.atlassian.com/secure/useravatar?avatarId=10605"
				},
				"displayName": "Bryan Rollins [Atlassian]",
				"active": "true"
			},
			"changelog": {
				"items": [
					{
						"toString": "A new summary.",
						"to": null,
						"fromString": "What is going on here?????",
						"from": null,
						"fieldtype": "jira",
						"field": "summary"
					},
					{
						"toString": "New Feature",
						"to": "2",
						"fromString": "Improvement",
						"from": "4",
						"fieldtype": "jira",
						"field": "issuetype"
					}
				],
				"id": 10124
			},
			"comment": {
				"self": "https://jira.atlassian.com/rest/api/2/issue/10148/comment/252789",
				"id": "252789",
				"author": {
					"self": "https://jira.atlassian.com/rest/api/2/user?username=brollins",
					"name": "brollins",
					"emailAddress": "bryansemail@atlassian.com",
					"avatarUrls": {
						"16x16": "https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
						"48x48": "https://jira.atlassian.com/secure/useravatar?avatarId=10605"
					},
					"displayName": "Bryan Rollins [Atlassian]",
					"active": true
				},
				"body": "Just in time for AtlasCamp!",
				"updateAuthor": {
					"self": "https://jira.atlassian.com/rest/api/2/user?username=brollins",
					"name": "brollins",
					"emailAddress": "brollins@atlassian.com",
					"avatarUrls": {
						"16x16": "https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
						"48x48": "https://jira.atlassian.com/secure/useravatar?avatarId=10605"
					},
					"displayName": "Bryan Rollins [Atlassian]",
					"active": true
				},
				"created": "2011-06-07T10:31:26.805-0500",
				"updated": "2011-06-07T10:31:26.805-0500"
			},
			"webhookEvent": "jira:issue_updated"
		}`,
		expected: SlackMessage{
			Text: "<@brollins> commented to the issue:",
			Attachments: SlackMessageAttachments{{
				Title:     "JRA-20002: I feel the need for speed",
				TitleLink: "https://jira.atlassian.com/browse/JRA-20002",
				Text:      "Just in time for AtlasCamp!",
				Timestamp: 1519952323,
			}},
		},
	},
}

func TestFormatJIRAEventToSlackMessage(t *testing.T) {
	for i := 0; i < len(fixtures); i++ {
		fixture := fixtures[i]
		expected := fixture.expected

		jiraEvent := new(JIRAEvent)
		if err := json.Unmarshal([]byte(fixture.eventJSON), jiraEvent); err != nil {
			t.Fatal(err)
		}
		actual := FormatJIRAEventToSlackMessage(jiraEvent, SlackAPIDialect("slack"))
		if actual.Text != expected.Text {
			t.Errorf("Text = %s, want %s", actual.Text, expected.Text)
		}
		if !reflect.DeepEqual(actual.Attachments, expected.Attachments) {
			t.Errorf("Attachments\n actual = %+v,\n expected = %+v", actual.Attachments, expected.Attachments)
		}
	}
}
