# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple Go web server built with the Echo framework. The project is currently minimal, containing only a basic "Hello, World!" HTTP server that listens on port 1323.

## Common Commands

**Run the application:**
```bash
go run cmd/server/main.go
```

**Build the application:**
```bash
go build -o bin/server cmd/server/main.go
```

**Install/update dependencies:**
```bash
go mod tidy
```

**Test the application:**
```bash
go test ./...
```

## Architecture

This project follows Go best practices with a layered architecture:

- **cmd/server/**: Application entry point
- **internal/**: Private application code
  - **handlers/**: HTTP request handlers (controllers)
  - **services/**: Business logic layer
  - **repository/**: Data access layer
  - **config/**: Configuration management
  - **middleware/**: Custom middleware
- **pkg/**: Reusable library code
  - **models/**: Data models and structures
  - **utils/**: Utility functions
- **docs/**: Project documentation

For detailed folder structure, see [docs/FOLDER_STRUCTURE.md](./docs/FOLDER_STRUCTURE.md)

## Development Notes

- Uses Echo v4 web framework for HTTP routing and middleware
- DynamoDB integration with AWS SDK v2 for data persistence
- Environment-based configuration for deployment flexibility
- Standard Go project structure with go.mod for dependency management

## Database Setup

- **Important**: DynamoDB table must be created before running the application
- See [DYNAMODB_SETUP.md](./DYNAMODB_SETUP.md) for complete setup instructions
- Required environment variables:
  - `DYNAMODB_TABLE_NAME` (default: "todos")
  - `AWS_REGION` (default: "us-east-1")
  - `PORT` (default: "1323")

## Local Development

**Prerequisites:**
- AWS CLI configured with appropriate credentials
- DynamoDB table created (see DYNAMODB_SETUP.md)

**Quick start:**
```bash
# Set environment variables
export DYNAMODB_TABLE_NAME=todos
export AWS_REGION=us-east-1

# Run application
go run cmd/server/main.go
```