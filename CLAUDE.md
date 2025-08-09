# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go commons package (`github.com/luciano-personal-org/commons`) that provides utility functions for file operations and data parsing. It's part of a larger algo-platform project structure.

## Development Commands

This is a Go package without a main function, so it's intended to be used as a library:

- **Build/Compile Check**: `go build .`
- **Test**: `go test .` (if tests are added)
- **Format Code**: `go fmt .`
- **Vet Code**: `go vet .`
- **Module Tidy**: `go mod tidy`

## Code Architecture

The package currently contains utility functions:

- `listFilesByPatternAndCreationTime`: Filters files by prefix pattern and sorts by modification time (descending)
- `parseFloat`: Simple string to float64 conversion utility

Key implementation notes:
- Uses `ioutil.ReadDir` (deprecated) - consider using `os.ReadDir` for new code
- The `listFilesByPatternAndCreationTime` function is currently unexported (lowercase) but provides file filtering and sorting functionality
- Error handling is minimal in `parseFloat` - errors are ignored

## Go Version

The project uses Go 1.23.2 as specified in go.mod.