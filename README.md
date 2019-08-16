# jira-to-slack [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack and Mattermost integration for notifying Jira events.
It is written in Go and ready on App Engine.


## Examples

### Slack

<img width="680" alt="jira-to-slack" src="https://user-images.githubusercontent.com/321266/36666061-c14e272e-1b2c-11e8-9e93-1f8f2857cbe0.png">


### Mattermost

<img width="638" alt="jira-to-mattermost" src="https://user-images.githubusercontent.com/321266/42192807-24339c98-7ea6-11e8-98b1-14b558c0d8bb.png">


## Getting Started

### 1. Setup Slack/Mattermost

Create an Incoming Webhook on your [Slack](https://my.slack.com/services/new/incoming-webhook) or [Mattermost](https://docs.mattermost.com/developer/webhooks-incoming.html).

### 2. Setup Jira

Create a [Webhook](https://developer.atlassian.com/server/jira/platform/webhooks/) on your Jira cloud or server.
Set the URL of the Webhook as follows:

```
https://jira-to-slack.appspot.com/?webhook=https://hooks.slack.com/YOUR_HOOK
```

You can add the following query parameters:

Name | Value | Default | Example
-----|-------|---------|--------
`webhook` | Slack/Mattermost Webhook URL | Mandatory | `https://hooks.slack.com/YOUR_HOOK`
`username` | BOT username | - | `JIRA`
`icon` | BOT icon emoji or URL | - | `:speech_baloon:` or `https://.../jira.png`
`dialect` | API dialect | `slack` | `slack` or `mattermost`
`debug` | Dump Jira and Slack messages to console | `0` | `0` or `1`

For example,

<img width="752" alt="jira-webhook-setup" src="https://user-images.githubusercontent.com/321266/42193983-a4c5fd32-7eac-11e8-979d-ae8103ae2672.png">

You can deploy jira-to-slack to your server as well.
See the later section for details.

### 3. Test notification

Create a ticket on your Jira and a message will be sent to your Slack/Mattermost.
You can turn on debug logs by setting the query parameter `debug=1`.


## Deploy to your server

### Standalone

Download the latest release and run the command:

```sh
./jira-to-slack
```

### Docker

Run the Docker image as follows:

```sh
docker run --rm -p 3000:3000 int128/jira-to-slack
```

### App Engine

You can deploy jira-to-slack to App Engine:

```sh
# Install SDK
brew cask install google-cloud-sdk
gcloud components install app-engine-go

# Run
dev_appserver.py appengine/app.yaml

# Deploy
gcloud app deploy --project=jira-to-slack appengine/app.yaml
```


## How it works

### Triggers

`jira-to-slack` sends a message to the Slack channel on the following triggers:

- Someone created an issue.
- Someone commented to an issue.
- Someone assigned an issue.
- Someone updated summary or description of an issue.
- Someone deleted an issue.

### Mentions

`jira-to-slack` sends mentions to reporter and assignee of the issue.

If the issue or comment has mentions (Slack style `@foo` or JIRA style `[~foo]`), `jira-to-slack` sends the mentions as well.


## Other solutions

[JIRA Mattermost Webhook Bridge](https://github.com/vrenjith/jira-matter-bridge). Great work. This is almost perfect but notifies many events so it may be noisy.

[Mattermost official JIRA Webhook Plugin](https://docs.mattermost.com/integrations/jira.html). This is still beta and in progress. Currently this does not notify comment.


## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to open your issues or pull requests.


### Development

Start the server:

```sh
make
./jira-to-slack
```

App Engine:

```sh
make run-appengine
```

### E2E Test

You can send actual payloads of actual Jira events by the following script:

```sh
# Slack
SLACK_WEBHOOK="https://hooks.slack.com/xxx&username=JIRA&icon=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0" ./pkg/formatter/testdata/post_jira_events.sh

# Mattermost
SLACK_WEBHOOK="https://mattermost.example.com/hooks/xxx&username=JIRA&icon=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0&dialect=mattermost" ./pkg/formatter/testdata/post_jira_events.sh
```
