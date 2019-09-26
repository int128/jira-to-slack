# jira-to-slack [![CircleCI](https://circleci.com/gh/int128/jira-to-slack.svg?style=shield)](https://circleci.com/gh/int128/jira-to-slack)

This is a Slack and Mattermost integration for notifying Jira events.
It is written in Go and ready on Docker, Cloud Run, App Engine and Lambda.


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
`channel` | Channel to show messages at | - | `some-public-channel`
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

It binds port 3000 by default.
You can set the port by `PORT` environment variable.

```sh
PORT=8080 ./jira-to-slack
```

### Docker / Docker Compose

Pull the image from [Docker Hub](https://hub.docker.com/r/int128/jira-to-slack) and run it as follows:

```sh
docker run --rm -p 3000:3000 int128/jira-to-slack:v1.9.0
```

You can run it on Docker Compose.

```yaml
services:
  jira-to-slack:
    image: int128/jira-to-slack:v1.9.0
    # Expose host port 8080
    ports: "8080:3000"
```

### Google Cloud Run

Click the button to deploy the image to Google Cloud Run.

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/jira-to-slack/jira-to-slack&cloudshell_git_repo=https://github.com/int128/jira-to-slack.git)

Alternatively can deploy the image `gcr.io/jira-to-slack/jira-to-slack` on Google Cloud Console.

### Google App Engine

You can deploy the application to Google App Engine.

```sh
# Install SDK
brew cask install google-cloud-sdk
gcloud components install app-engine-go

# Run locally
make -C appengine run

# Deploy to cloud
gcloud app deploy --project=jira-to-slack appengine/app.yaml
```

### AWS Lambda

You can deploy the application to AWS Lambda.

```sh
# Run locally
make -C lambda run

# Deploy to cloud
make -C lambda deploy SAM_S3_BUCKET_NAME=YOUR_BUCKET_NAME
```

You need to create a S3 bucket in the same region before deploying.


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

### E2E Test

You can send actual payloads of actual Jira events by the following script:

```sh
# Slack
SLACK_WEBHOOK="https://hooks.slack.com/xxx&username=JIRA&icon=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0" ./pkg/formatter/testdata/post_jira_events.sh

# Mattermost
SLACK_WEBHOOK="https://mattermost.example.com/hooks/xxx&username=JIRA&icon=https://lh3.googleusercontent.com/GkgChJMixx9JAmoUi1majtfpjg1Ra86gZR0GCehJfVcOGQI7Ict_TVafXCtJniVn3R0&dialect=mattermost" ./pkg/formatter/testdata/post_jira_events.sh
```
