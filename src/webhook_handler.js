const Slack = require('./slack');
const WebhookMessage = require('./webhook_message');

const slack = new Slack(
  process.env.SLACK_WEBHOOK, {
    username: process.env.SLACK_USERNAME,
    iconEmoji: process.env.SLACK_ICON_EMOJI,
    iconUrl: process.env.SLACK_ICON_URL,
  });

module.exports = async req => {
  if (typeof req !== 'object') {
    throw new TypeError(`Request must be a valid object: ${req}`);
  }
  const message = new WebhookMessage(req.body);
  if (message.isValid()) {
    return await slack.send({
      attachments: [
        {
          title: message.getTitle(),
          title_link: message.getTitleLink(),
          pretext: message.getPretext(),
          text: message.getText(),
          ts: message.getUpdatedTimestamp(),
        },
      ],
    });
  }
}
