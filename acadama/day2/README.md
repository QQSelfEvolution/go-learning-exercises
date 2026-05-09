# JSON Formatter & Log Analyzer

A Go CLI tool for JSON formatting and log analysis.

## Features

- Read JSON from file or stdin
- Pretty print with colors and indentation
- Filter and search JSON keys/values
- Log analysis mode

## Installation

```bash
go install ./...
```

## Usage

### Format JSON from file
```bash
go run main.go -f input.json
```

### Format JSON from stdin
```bash
cat input.json | go run main.go
```

### Pretty print with colors
```bash
go run main.go -f input.json -p
```

### Search for specific key
```bash
go run main.go -f input.json -k "name"
```

### Filter by value
```bash
go run main.go -f input.json -v "John"
```

### Minify JSON
```bash
go run main.go -f input.json -m
```

## Options

- `-f, --file`: Input JSON file path
- `-p, --pretty`: Pretty print with colors
- `-k, --key`: Filter by key name
- `-v, --value`: Filter by value
- `-m, --minify`: Minify JSON output
- `-c, --color`: Enable colored output (default: true)

## Examples

```bash
# Pretty print with colors
echo '{"name":"John","age":30}' | go run main.go -p

# Search for specific key
echo '{"users":[{"name":"John"},{"name":"Jane"}]}' | go run main.go -k "name"

# Minify JSON
go run main.go -f data.json -m
```
