# JIRA to Slack

This is an integration for notification from JIRA to Slack, running on Node.js and Docker.

## Getting Started

Node.js:

```bash
npm install
export SLACK_WEBHOOK=https://hooks.slack.com/...
npm start
```

Docker:

```bash
docker run -e SLACK_WEBHOOK=https://hooks.slack.com/... -p 3000:3000 int128/jira-to-slack
```

Kubernetes:

```bash
kubectl apply -f kubernetes.yaml
```

Then, create a webhook on your JIRA and point to `http://your-node-js-server:3000`.

### Settings

You can set the following environment variables:

- `SLACK_WEBHOOK` - Webhook URL to the Slack channel (Required)
- `SLACK_USERNAME` - Username of the BOT (Optional)
- `SLACK_ICON_EMOJI` - Icon emoji of the BOT (Optional)
- `SLACK_ICON_URL` - Icon URL of the BOT (Optional)

## Contribution

This is an open source software licensed under Apache License 2.0.
Feel free to book your issues or pull requests.

### See Also

- https://developer.atlassian.com/server/jira/platform/webhooks/
