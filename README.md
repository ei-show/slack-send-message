# slack-send-message

[![CI](https://github.com/ei-show/slack-send-message/workflows/CI/badge.svg)](https://github.com/ei-show/slack-send-message/actions?query=workflow%3ACI)
[![Release](https://github.com/ei-show/slack-send-message/workflows/Release/badge.svg)](https://github.com/ei-show/slack-send-message/actions?query=workflow%3ARelease)

A CLI tool written in Go to send messages to Slack channels using the Slack API. Designed for use in CI/CD pipelines.

## Features

- Send messages to Slack channels
- Reply to existing threads
- Environment variable support for API keys
- Returns thread timestamp for further operations

## Installation

### Download from releases

Download the latest binary for your platform from the [releases page](https://github.com/ei-show/slack-send-message/releases).

### Build from source

```bash
go build -o slack-send-message .
```

## Development

### Running tests

```bash
go test -v ./...
```

### Running with coverage

```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

## Usage

### Send a message to a channel

**Recommended for production/CI/CD**: Use environment variable for API key (more secure):

```bash
export SLACK_API_KEY=xoxb-your-token
slack-send-message --channel-id C123456789 "Hello, Slack!"
```

**For local testing only**: You can pass the API key via flag (not recommended for production/CI/CD due to security concerns):

```bash
slack-send-message --api-key xoxb-your-token --channel-id C123456789 "Hello, Slack!"
```

The command will output the thread timestamp:

```
1234567890.123456
```

### Reply to a thread

**Recommended for production/CI/CD**: Use environment variable for API key (more secure):

```bash
export SLACK_API_KEY=xoxb-your-token
slack-send-message reply --channel-id C123456789 --thread-ts 1234567890.123456 "This is a reply"
```

**For local testing only**: You can pass the API key via flag (not recommended for production/CI/CD due to security concerns):

```bash
slack-send-message reply --api-key xoxb-your-token --channel-id C123456789 --thread-ts 1234567890.123456 "This is a reply"
```

## Parameters

### Main command (send message)

- `--api-key`: Slack API key (can also use `SLACK_API_KEY` environment variable, parameter takes priority)
- `--channel-id`: Slack channel ID (required)
- `[message]`: Message to send (required, positional argument)

### Reply subcommand

- `--api-key`: Slack API key (can also use `SLACK_API_KEY` environment variable, parameter takes priority)
- `--channel-id`: Slack channel ID (required)
- `--thread-ts`: Thread timestamp to reply to (required)
- `[message]`: Message to send as reply (required, positional argument)

## CI/CD Usage Example

```yaml
# GitHub Actions example
- name: Send Slack notification
  env:
    SLACK_API_KEY: ${{ secrets.SLACK_API_KEY }}
  run: |
    THREAD_TS=$(slack-send-message --channel-id C123456789 "Build started")
    echo "Thread TS: $THREAD_TS"
    
    # Later, reply to the thread
    slack-send-message reply --channel-id C123456789 --thread-ts $THREAD_TS "Build completed successfully"
```
