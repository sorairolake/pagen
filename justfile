# SPDX-FileCopyrightText: 2025 Shun Sakai
#
# SPDX-License-Identifier: CC0-1.0

alias build := build-debug
alias fmt := golangci-lint-fmt
alias lint := golangci-lint-run

# Run default recipe
_default:
    just -l

# Build `pagen` command in debug mode
build-debug $CGO_ENABLED="0":
    go build

# Build `pagen` command in release mode
build-release $CGO_ENABLED="0":
    go build -ldflags="-s -w" -trimpath

# Remove generated artifacts
clean:
    go clean

# Run tests
test:
    go test ./...

# Run `golangci-lint`
golangci-lint: golangci-lint-fmt golangci-lint-run

# Run the formatter
golangci-lint-fmt:
    golangci-lint fmt

# Run the linter
golangci-lint-run:
    golangci-lint run

# Run the linter for GitHub Actions workflow files
lint-github-actions:
    actionlint -verbose

# Run the formatter for the README
fmt-readme:
    npx prettier -w README.md

# Increment the version
bump part:
    bump-my-version bump {{ part }}
