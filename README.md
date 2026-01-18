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

## Release Process

This project uses an automated release workflow powered by [Release Please](https://github.com/googleapis/release-please), which handles changelog generation, version bumping, and binary releases.

### How It Works

The release workflow (`.github/workflows/release-please.yml`) consists of two jobs:

#### 1. Release Please Job

- **Triggers**: Automatically runs on every push to the `main` branch, or can be manually triggered
- **Purpose**: Creates and maintains a release PR that automatically updates the changelog and version
- **How it works**:
  - Scans commit messages following [Conventional Commits](https://www.conventionalcommits.org/) format
  - Automatically generates a changelog based on commits since the last release
  - Creates or updates a release PR with version bump and changelog
  - When the release PR is merged, automatically creates a GitHub release with a git tag

#### 2. Build and Release Job

- **Triggers**: Runs only when a new release is created (after merging the release PR)
- **Purpose**: Builds and uploads binaries for multiple platforms
- **Steps**:
  1. Checks out the code
  2. Sets up Go using the version specified in `go.mod`
  3. Runs all tests to ensure quality
  4. Builds binaries for 6 different platform combinations:
     - Linux (amd64, arm64)
     - macOS/Darwin (amd64, arm64)
     - Windows (amd64, arm64)
  5. Uploads all built binaries as release assets to the GitHub release

### For Contributors

When making changes, use conventional commit messages to ensure proper changelog generation:

- `feat:` - New features (triggers minor version bump)
- `fix:` - Bug fixes (triggers patch version bump)
- `docs:` - Documentation changes
- `chore:` - Maintenance tasks
- `BREAKING CHANGE:` - Breaking changes (triggers major version bump)

Example:
```bash
git commit -m "feat: add support for message attachments"
git commit -m "fix: handle connection timeout errors"
```

The Release Please bot will automatically create or update a release PR, grouping all changes since the last release. Once this PR is merged, a new release with binaries will be automatically published.
