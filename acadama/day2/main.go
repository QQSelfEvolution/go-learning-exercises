package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Color codes for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

var (
	filePath  = flag.String("f", "", "Input JSON file path")
	pretty   = flag.Bool("p", false, "Pretty print with colors")
	key      = flag.String("k", "", "Filter by key name")
	value    = flag.String("v", "", "Filter by value")
	minify   = flag.Bool("m", false, "Minify JSON output")
	noColor  = flag.Bool("nc", false, "Disable colored output")
)

func main() {
	flag.Parse()

	var jsonData []byte
	var err error

	// Read JSON from file or stdin
	if *filePath != "" {
		jsonData, err = os.ReadFile(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Read from stdin
		jsonData, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	}

	// Trim whitespace
	jsonData = strings.TrimSpace(jsonData)

	// Parse JSON
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Apply filters
	if *key != "" {
		data = filterByKey(data, *key)
	}
	if *value != "" {
		data = filterByValue(data, *value)
	}

	// Output JSON
	var output []byte
	if *minify {
		output, err = json.Marshal(data)
	} else {
		if *pretty {
			output, err = json.MarshalIndent(data, "", "  ")
		} else {
			output, err = json.MarshalIndent(data, "", "  ")
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Print output
	if *pretty && !*noColor {
		printColored(output)
	} else {
		fmt.Println(string(output))
	}
}

func filterByKey(data interface{}, key string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, val := range v {
			if strings.Contains(strings.ToLower(k), strings.ToLower(key)) {
				result[k] = val
			}
			result[k] = filterByKey(val, key)
		}
		return result
	case []interface{}:
		result := make([]interface{}, 0)
		for _, item := range v {
			filtered := filterByKey(item, key)
			if filtered != nil {
				result = append(result, filtered)
			}
		}
		return result
	default:
		return nil
	}
}

func filterByValue(data interface{}, value string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, val := range v {
			result[k] = filterByValue(val, value)
			if result[k] != nil {
				if strVal, ok := val.(string); ok && strings.Contains(strings.ToLower(strVal), strings.ToLower(value)) {
					result[k] = val
				} else if result[k] != nil {
					result[k] = val
				}
			}
		}
		// Check if map has any non-nil values
		for _, val := range result {
			if val != nil {
				return result
			}
		}
		return nil
	case []interface{}:
		result := make([]interface{}, 0)
		for _, item := range v {
			filtered := filterByValue(item, value)
			if filtered != nil {
				result = append(result, filtered)
			}
		}
		if len(result) > 0 {
			return result
		}
		return nil
	case string:
		if strings.Contains(strings.ToLower(v), strings.ToLower(value)) {
			return data
		}
		return nil
	default:
		return nil
	}
}

func printColored(json []byte) {
	colorEnabled := !*noColor

	scanner := bufio.NewScanner(strings.NewReader(string(json)))
	for scanner.Scan() {
		line := scanner.Text()
		coloredLine := colorizeJSON(line, colorEnabled)
		fmt.Println(coloredLine)
	}
}

func colorizeJSON(line string, enabled bool) string {
	if !enabled {
		return line
	}

	// Color keys
	line = strings.ReplaceAll(line, `"`, ColorCyan+`"`+ColorReset)

	// Find and color string values
	inString := false
	result := []byte{}
	for i := 0; i < len(line); i++ {
		if line[i] == '"' && (i == 0 || line[i-1] != '\\') {
			inString = !inString
			if inString {
				result = append(result, ColorGreen+'"'...)
			} else {
				result = append(result, '"'+ColorReset...)
			}
		} else if inString {
			result = append(result, line[i])
		} else {
			result = append(result, line[i])
		}
	}

	return string(result)
}
