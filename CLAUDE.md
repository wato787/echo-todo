# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple Go web server built with the Echo framework. The project is currently minimal, containing only a basic "Hello, World!" HTTP server that listens on port 1323.

## Common Commands

**Run the application:**
```bash
go run main.go
```

**Build the application:**
```bash
go build -o echo-todo
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

- **main.go**: Entry point containing the Echo server setup and a single GET route handler
- **go.mod**: Module definition with Echo v4 framework dependency
- The server runs on port 1323 and currently serves a single endpoint at "/"

## Development Notes

- Uses Echo v4 web framework for HTTP routing and middleware
- Standard Go project structure with go.mod for dependency management
- No database or external services currently configured
- No tests or additional middleware implemented yet