Introduction
=============

This project contains codes to validate [Walter](https://github.com/walter-cd/walter)

Environment Variables
======================

You need to set the following envrionment variables before running the Walter integration tests.

| Variable          | Description       |
|:------------------|:------------------|
| HIPCHAT_ROOM_ID   | Your HipChat room name |
| HIPCHAT_USER_NAME | Your HipChat user name |
| HIPCHAT_TOKEN     | Your HipChat token ([API v1](https://www.hipchat.com/docs/api)) |
| HIPCHAT_TWO_TOKEN | Your HipChat token ([API v2](https://www.hipchat.com/docs/apiv2)) |
| SLACK_CHANNEL     | Your Slack channel name |
| SLACK_USER_NAME   | Your Slack user name  |
| SLACK_URL         | Your Slack access url (see https://my.slack.com/services/new/incoming-webhook)|
| GITHUB_TOKEN      | Your [GitHub token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) |

Running tests
===================

```
$ git@github.com:takahi-i/walter-integration-test.git
$ cd walter-integration-test
$ sh test.sh
```
