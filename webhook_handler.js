const Slack = require('./slack');
const WebhookMessage = require('./webhook_message');

const slack = new Slack(
  process.env.SLACK_WEBHOOK, {
    username: process.env.SLACK_USERNAME,
    iconEmoji: process.env.SLACK_ICON_EMOJI,
    iconUrl: process.env.SLACK_ICON_URL,
  });

module.exports = async function (req) {
  if (typeof req !== 'object') {
    throw new TypeError(`Request must be a valid object: ${req}`);
  }
  const message = new WebhookMessage(req.body);
  const text = message.formatText();
  if (text) {
    await slack.send(text);
  }
}
