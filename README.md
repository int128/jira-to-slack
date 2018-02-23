# JIRA to Slack integration [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This notifies JIRA events to a Slack channel or a Mattermost channel.

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

Run a container.

```bash
# Docker
docker run --rm -p 3000:3000 \
  -e SLACK_WEBHOOK=https://hooks.slack.com/... \
  -e SLACK_USERNAME=JIRA \
  -e SLACK_ICON_URL=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0 \
  int128/jira-to-slack

# Kubernetes
kubectl apply -f kubernetes.yaml
```

### 3. Create JIRA webhook

Create a webhook on your JIRA.
See also the [JIRA Server Webhooks](https://developer.atlassian.com/server/jira/platform/webhooks/) for details.

## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.
