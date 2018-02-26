const JIRA_FIELDS_TO_NOTIFY_UPDATE = ['summary', 'description', 'assignee'];

module.exports = class WebhookMessage {
  /**
   * Constructor.
   * @param {object} body JSON payload
   */
  constructor(body, { dialect }) {
    if (typeof body !== 'object') {
      throw new TypeError(`Request body must be a valid object: ${body}`);
    }
    this._body = body;
    this._dialect = dialect;
  }

  /**
   * @returns {boolean} true if the message is valid and to be sent
   */
  isValid() {
    const { webhookEvent, changelog, comment } = this._body;
    switch (webhookEvent) {
      case 'jira:issue_updated':
        if (comment) {
          return true;
        }
        if (changelog) {
          return changelog.items.find(item =>
            JIRA_FIELDS_TO_NOTIFY_UPDATE.find(field => item.field === field));
        }
        break;
      case 'jira:issue_created':
        return true;
      case 'jira:issue_deleted':
        return true;
    }
  }

  /**
   * @returns {string} title of the message
   */
  getTitle() {
    const { issue } = this._body;
    return `${issue.key}: ${issue.fields.summary}`;
  }

  /**
   * @returns {string} link URL of the message
   */
  getTitleLink() {
    const { issue } = this._body;
    return `${issue.self.replace(/\/rest\/api\/.+/, '')}/browse/${issue.key}`;
  }

  /**
   * @returns {string} pretext of the message (who)
   */
  getPretext() {
    const { webhookEvent, user, comment, issue } = this._body;
    const username = this._formatUsername(user);
    let verb;
    switch (webhookEvent) {
      case 'jira:issue_updated':
        if (comment) {
          verb = 'commented to';
        } else {
          verb = 'updated';
        }
        break;
      case 'jira:issue_created':
        verb = 'created';
        break;
      case 'jira:issue_deleted':
        verb = 'deleted';
        break;
    }
    let assignee = '';
    if (issue.fields.assignee) {
      assignee = `(assigned to ${this._formatUsername(issue.fields.assignee)}>)`;
    }
    return `${username} ${verb} the issue: ${assignee}`
  }

  /**
   * @returns {string} pretext of the message (what)
   */
  getText() {
    const { webhookEvent, issue, comment } = this._body;
    switch (webhookEvent) {
      case 'jira:issue_updated':
        if (comment) {
          return comment.body;
        } else {
          return issue.fields.description;
        }
      case 'jira:issue_created':
        return issue.fields.description;
      case 'jira:issue_deleted':
        return issue.fields.description;
    }
  }

  getUpdatedTimestamp() {
    const { timestamp } = this._body;
    if (typeof timestamp === 'number') {
      return timestamp / 1000;
    }
  }

  /**
   * Format the username.
   * @param {object} user user object that has user.name
   */
  _formatUsername(user) {
    switch (this._dialect) {
      case 'mattermost':
        return `@${user.name}`;
      default:
        return `<@${user.name}>`;
    }
  }
}
