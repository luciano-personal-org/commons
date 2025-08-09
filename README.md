# Commons

A Go utility package providing common file operations and data parsing functions.

## Installation

```bash
go get github.com/luciano-personal-org/commons
```

## Functions

### File Operations

- `listFilesByPatternAndCreationTime(directoryPath, searchPattern string) ([]os.FileInfo, error)`
  - Filters files by prefix pattern and sorts by modification time (descending)
  - Currently unexported (private function)

### Data Parsing

- `parseFloat(value string) float64`
  - Simple string to float64 conversion utility
  - Returns 0 if parsing fails

## Requirements

- Go 1.23.2 or later

## Development

```bash
# Compile check
go build .

# Format code
go fmt .

# Vet code
go vet .

# Tidy modules
go mod tidy
```