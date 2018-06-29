# JIRA to Slack Integration [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack integration for notifying JIRA events to a channel.
It supports Mattermost as well.

<img width="680" alt="jira-to-slack" src="https://user-images.githubusercontent.com/321266/36666061-c14e272e-1b2c-11e8-9e93-1f8f2857cbe0.png">

## Getting Started

Create an [incoming webhook](https://my.slack.com/services/new/incoming-webhook) on your Slack team.

Run a server.

```sh
./jira-to-slack
```

Create a [webhook](https://developer.atlassian.com/server/jira/platform/webhooks/) on your JIRA server.
You can add the following query parameters to the webhook URL.

Name | Value | Example value
-----|-------|--------------
`webhook` | Slack webhook URL (Mandatory) | `https://hooks.slack.com/xxx`
`username` | Username of the BOT | `JIRA`
`icon` | Icon emoji or URL of the BOT | `:speech_baloon:` or `http://.../jira.png`
`dialect` | Slack API dialect (Default to `slack`) | `slack` or `mattermost`

For example:

```
https://jira-to-slack.example.com/?webhook=https://hooks.slack.com/xxx&username=JIRA
```


### Using Docker

```bash
docker run --rm -p 3000:3000 int128/jira-to-slack
```

### Using Kubernetes

You can install the Helm chart from https://github.com/int128/devops-kompose/tree/master/jira-to-slack.


## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.


### Development

Start the server:

```sh
go build && ./jira-to-slack
```

### E2E Test

You can send actual payloads of actual JIRA events by the following script:

```sh
SLACK_WEBHOOK=https://hooks.slack.com/xxx ./testdata/post_jira_events.sh
```
