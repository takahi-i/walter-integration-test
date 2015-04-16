Introduction
=============

This project contains codes to validate [Walter](https://github.com/walter-cd/walter)

Environment Variables
======================

You need to set the following envrionment variables before running the Walter integration tests.

| Variable          | Description       |
|:------------------|:------------------|
| HIPCHAT_ROOM_ID   | HipChat room name |
| HIPCHAT_USER_NAME | HipChat user name |
| HIPCHAT_TOKEN     | HipChat token (API v1) |
| HIPCHAT_TWO_TOKEN | HipChat token (API v2) |
| SLACK_CHANNEL     | Slack channel name |
| SLACK_USER_NAME   | Slack user name  |
| SLACK_URL         | Slack access url |
| GITHUB_TOKEN      | GitHub Token |
| GITHUB_REPOSITORY | GitHub repository name |
| GITHUB_GROUP      | GitHub group name |

Running tests
===================

    $ git@github.com:takahi-i/walter-integration-test.git
    $ cd walter-integration-test
    $ sh test.sh
