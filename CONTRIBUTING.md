# Contributing to slack-send-message

Thank you for your interest in contributing to slack-send-message! This document provides guidelines for contributing to this project.

## Development

### Running tests

```bash
go test -v ./...
```

### Running with coverage

```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

### Building from source

```bash
go build -o slack-send-message .
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

## Commit Message Guidelines

When making changes, use conventional commit messages to ensure proper changelog generation:

- `feat:` - New features (triggers minor version bump)
- `fix:` - Bug fixes (triggers patch version bump)
- `docs:` - Documentation changes
- `chore:` - Maintenance tasks
- `BREAKING CHANGE:` - Breaking changes (triggers major version bump)

### Examples

```bash
git commit -m "feat: add support for message attachments"
git commit -m "fix: handle connection timeout errors"
git commit -m "docs: update installation instructions"
```

The Release Please bot will automatically create or update a release PR, grouping all changes since the last release. Once this PR is merged, a new release with binaries will be automatically published.
