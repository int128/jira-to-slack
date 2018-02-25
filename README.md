# JIRA to Slack Integration [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack integration for notifying JIRA events to a channel.
It supports Mattermost as well.

<img width="700" alt="jita-to-slack" src="https://user-images.githubusercontent.com/321266/36641377-14af7532-1a72-11e8-95b8-a5b12b3fa3e0.png">

## Getting Started

### Prerequisite

Create [an incoming webhook](https://my.slack.com/services/new/incoming-webhook) on your Slack team.

### Using Docker

Run a container.

```bash
docker run --rm -p 3000:3000 \
  -e SLACK_WEBHOOK=https://hooks.slack.com/... \
  -e SLACK_USERNAME=JIRA \
  -e SLACK_ICON_URL=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0 \
  int128/jira-to-slack
```

And then, create [a webhook](https://developer.atlassian.com/server/jira/platform/webhooks/) on your JIRA server.

### Using Kubernetes

You can install the [Helm](https://github.com/kubernetes/helm) chart as follows:

```yaml
# jira-to-slack.yaml
slack:
  ## Slack webhook URL (mandatory)
  webhook: https://hooks.slack.com/...
  ## Slack username
  # username: JIRA
  ## Slack icon emoji
  # iconEmoji: ":speech_baloon:"
  ## Slack icon image URL
  #iconImageURL: https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0
```

```bash
helm repo add jira-to-slack https://int128.github.io/jira-to-slack
helm repo update
helm install jira-to-slack/jira-to-slack -f jira-to-slack.yaml
```

And then, create [a webhook](https://developer.atlassian.com/server/jira/platform/webhooks/) on your JIRA server.
If both the JIRA server and jira-to-slack are running in the same namespace, point the webhook to the service URL like `http://jira-to-slack-jira-to-slack`.

## Customize

### Slack

You can set the BOT username and icon in the Slack webhook settings.
Instead you can set the following environment variables to the container:

Name | Value | Example
-----|-------|--------
`SLACK_WEBHOOK` | Slack webhook URL (Required) | `https://hooks.slack.com/...`
`SLACK_USERNAME` | Username of the BOT (Optional) | `JIRA`
`SLACK_ICON_EMOJI` | Icon emoji of the BOT (Optional) | `:speech_baloon:`
`SLACK_ICON_URL` | Icon URL of the BOT (Optional) | `http://.../jira.png`

### JIRA

You can filter projects by JQL in the webhook settings of the JIRA server.

## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.
