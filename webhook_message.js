const JIRA_FIELDS_TO_NOTIFY_UPDATE = ['summary', 'description', 'assignee'];

module.exports = class WebhookMessage {
  constructor(body) {
    if (typeof body !== 'object') {
      throw new TypeError(`Request body must be a valid object: ${body}`);
    }
    this._body = body;
  }

  formatText() {
    const { webhookEvent } = this._body;
    switch (webhookEvent) {
      case 'jira:issue_updated': return this.formatIssueUpdated();
      case 'jira:issue_created': return this.formatIssueCreated();
      case 'jira:issue_deleted': return this.formatIssueDeleted();
    }
  }

  formatIssueUpdated() {
    const { user, issue, changelog, comment } = this._body;
    const isChanged = changelog && changelog.items.find(item => JIRA_FIELDS_TO_NOTIFY_UPDATE.find(field => item.field === field));
    if (isChanged) {
      return `${_user(user)} updated: ${_issue(issue)}`;
    }
    if (comment) {
      return `${_user(user)} commented: ${_comment(issue, comment)}`;
    }
  }

  formatIssueCreated() {
    const { user, issue } = this._body;
    return `${_user(user)} created: ${_issue(issue)}`;
  }

  formatIssueDeleted() {
    const { user, issue } = this._body;
    return `${_user(user)} deleted: ${_issue(issue)}`;
  }
}

const _user = user => `@${user.name}`;

const _issue = issue => _quote(
  `**[${issue.key}](${_issueUrl(issue)}) ${issue.fields.summary}** ${_assignee(issue)}\r\n${issue.fields.description || ''}`);

const _comment = (issue, comment) => _quote(
  `**[${issue.key}](${_issueUrl(issue)}) ${issue.fields.summary}** ${_assignee(issue)}\r\n${comment.body}`);

const _assignee = issue => issue.fields.assignee ? `(assigned to ${_user(issue.fields.assignee)})` : '';

const _issueUrl = issue => `${issue.self.replace(/\/rest\/api\/.+/, '')}/browse/${issue.key}`;

const _quote = content => content.replace(/^|\r\n|\r|\n/g, '\r\n> ');
