# slack-send-message

A CLI tool written in Go to send messages to Slack channels using the Slack API. Designed for use in CI/CD pipelines.

## Features

- Send messages to Slack channels
- Reply to existing threads
- Environment variable support for API keys
- Returns thread timestamp for further operations

## Installation

### Build from source

```bash
go build -o slack-send-message .
```

## Usage

### Send a message to a channel

```bash
slack-send-message --api-key xoxb-your-token --channel-id C123456789 --message "Hello, Slack!"
```

Using environment variable for API key:

```bash
export SLACK_API_KEY=xoxb-your-token
slack-send-message --channel-id C123456789 --message "Hello, Slack!"
```

The command will output the thread timestamp:

```
1234567890.123456
```

### Reply to a thread

```bash
slack-send-message reply --api-key xoxb-your-token --channel-id C123456789 --thread-ts 1234567890.123456 --message "This is a reply"
```

Using environment variable for API key:

```bash
export SLACK_API_KEY=xoxb-your-token
slack-send-message reply --channel-id C123456789 --thread-ts 1234567890.123456 --message "This is a reply"
```

## Parameters

### Main command (send message)

- `--api-key`: Slack API key (can also use `SLACK_API_KEY` environment variable, parameter takes priority)
- `--channel-id`: Slack channel ID (required)
- `--message`: Message to send (required)

### Reply subcommand

- `--api-key`: Slack API key (can also use `SLACK_API_KEY` environment variable, parameter takes priority)
- `--channel-id`: Slack channel ID (required)
- `--thread-ts`: Thread timestamp to reply to (required)
- `--message`: Message to send as reply (required)

## CI/CD Usage Example

```yaml
# GitHub Actions example
- name: Send Slack notification
  env:
    SLACK_API_KEY: ${{ secrets.SLACK_API_KEY }}
  run: |
    THREAD_TS=$(slack-send-message --channel-id C123456789 --message "Build started")
    echo "Thread TS: $THREAD_TS"
    
    # Later, reply to the thread
    slack-send-message reply --channel-id C123456789 --thread-ts $THREAD_TS --message "Build completed successfully"
```
