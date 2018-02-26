const JIRA_FIELDS_TO_NOTIFY_UPDATE = ['summary', 'description', 'assignee'];

module.exports = class WebhookMessage {
  constructor(body) {
    if (typeof body !== 'object') {
      throw new TypeError(`Request body must be a valid object: ${body}`);
    }
    this._body = body;
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
    let assignee = '';
    if (issue.fields.assignee) {
      assignee = `(assigned to @${issue.fields.assignee.name})`;
    }
    switch (webhookEvent) {
      case 'jira:issue_updated':
        if (comment) {
          return `@${user.name} commented: ${assignee}`;
        } else {
          return `@${user.name} updated: ${assignee}`;
        }
      case 'jira:issue_created':
        return `@${user.name} created: ${assignee}`;
      case 'jira:issue_deleted':
        return `@${user.name} deleted:`;
    }
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
    return parseInt(timestamp) / 1000;
  }
}
