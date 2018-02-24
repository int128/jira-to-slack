# JIRA to Slack Integration [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack integration for notifying JIRA events to a channel.
It supports Mattermost as well.

## Getting Started

### 1. Create Slack webhook

Create [an incoming webhook integration](https://my.slack.com/services/new/incoming-webhook) and get the webhook URL.

### 2. Run server

Set the following environment variables:

Name | Value | Example
-----|-------|--------
`SLACK_WEBHOOK` | Slack webhook URL (Required) | `https://hooks.slack.com/...`
`SLACK_USERNAME` | Username of the BOT (Optional) | `JIRA`
`SLACK_ICON_EMOJI` | Icon emoji of the BOT (Optional) | `:speech_baloon:`
`SLACK_ICON_URL` | Icon URL of the BOT (Optional) | `http://.../jira.png`

#### Using Docker

Run a container.

```bash
docker run --rm -p 3000:3000 \
  -e SLACK_WEBHOOK=https://hooks.slack.com/... \
  -e SLACK_USERNAME=JIRA \
  -e SLACK_ICON_URL=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0 \
  int128/jira-to-slack
```

#### Using Kubernetes

You can install the [Helm](https://github.com/kubernetes/helm) chart as follows:

```yaml
# jira-to-slack.yaml
slack:
  # Slack webhook URL (mandatory)
  webhook: https://hooks.slack.com/...
  # Slack username
  username: JIRA
  # Slack icon emoji
  iconEmoji: ":speech_baloon:"
  # Slack icon image URL
  iconImageURL: https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0
```

```bash
helm repo add jira-to-slack https://int128.github.io/jira-to-slack
helm repo update
helm install jira-to-slack/jira-to-slack -f jira-to-slack.yaml
```

### 3. Create JIRA webhook

Create a webhook on your JIRA.
See also the [JIRA Server Webhooks](https://developer.atlassian.com/server/jira/platform/webhooks/) for details.

## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.
