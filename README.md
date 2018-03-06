# JIRA to Slack Integration [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack integration for notifying JIRA events to a channel.
It supports Mattermost as well.

<img width="680" alt="jira-to-slack" src="https://user-images.githubusercontent.com/321266/36666061-c14e272e-1b2c-11e8-9e93-1f8f2857cbe0.png">

## Getting Started

### Prerequisite

Create [an incoming webhook](https://my.slack.com/services/new/incoming-webhook) on your Slack team.

### Using Docker

Run a container.

```bash
docker run --rm -p 3000:3000 \
  -e SLACK_WEBHOOK=https://hooks.slack.com/... \
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
  # iconImageURL: https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0
  ## Slack API dialect
  # dialect: slack
```

```bash
helm repo add int128 https://int128.github.io/helm-charts
helm repo update
helm install int128/jira-to-slack -f jira-to-slack.yaml
```

And then, create [a webhook](https://developer.atlassian.com/server/jira/platform/webhooks/) on your JIRA server.
If both the JIRA server and jira-to-slack are running in the same namespace, point the webhook to the service URL like `http://jira-to-slack-jira-to-slack`.

## Customize

### Slack

You can set the BOT username and icon in the Slack webhook settings.
Instead you can set the following environment variables to the container:

Name | Value | Example value
-----|-------|--------------
`SLACK_WEBHOOK` | Slack webhook URL (Required) | `https://hooks.slack.com/...`
`SLACK_USERNAME` | Username of the BOT (Optional) | `JIRA`
`SLACK_ICON_EMOJI` | Icon emoji of the BOT (Optional) | `:speech_baloon:`
`SLACK_ICON_URL` | Icon URL of the BOT (Optional) | `http://.../jira.png`
`SLACK_API_DIALECT` | Slack API dialect, defaults to `slack` | `slack` or `mattermost`

### JIRA

You can filter projects by JQL in the webhook settings of the JIRA server.

## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.

### Development

Start the server:

```sh
go build && SLACK_WEBHOOK=https://hooks.slack.com/... ./jira-to-slack
```

### E2E Test

You can send actual payloads of actual JIRA events by the following script:

```sh
./fixtures/post_jira_events.sh
```
