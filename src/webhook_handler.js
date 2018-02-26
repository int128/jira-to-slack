const Slack = require('./slack');
const WebhookMessage = require('./webhook_message');

const slack = new Slack(
  process.env.SLACK_WEBHOOK, {
    username: process.env.SLACK_USERNAME,
    iconEmoji: process.env.SLACK_ICON_EMOJI,
    iconUrl: process.env.SLACK_ICON_URL,
  });

const dialect = process.env.SLACK_API_DIALECT;

module.exports = async req => {
  if (typeof req !== 'object') {
    throw new TypeError(`Request must be a valid object: ${req}`);
  }
  const message = new WebhookMessage(req.body, { dialect });
  if (message.isValid()) {
    return await slack.send({
      // Use text instead of pretext in attachment,
      // due to the issue that @username in pretext is not linked on Mattermost.
      text: message.getPretext(),
      attachments: [
        {
          title: message.getTitle(),
          title_link: message.getTitleLink(),
          text: message.getText(),
          ts: message.getUpdatedTimestamp(),
        },
      ],
    });
  }
}
