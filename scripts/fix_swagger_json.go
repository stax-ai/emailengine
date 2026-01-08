package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: fix_swagger_json <input.json> <output.json>")
		os.Exit(2)
	}
	in := os.Args[1]
	out := os.Args[2]

	data, err := os.ReadFile(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read:", err)
		os.Exit(1)
	}
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal:", err)
		os.Exit(1)
	}
	clean := stripExamples(v)
	buf, err := json.MarshalIndent(clean, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "marshal:", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "mkdir:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(out, buf, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "write:", err)
		os.Exit(1)
	}
}

var backslashLetter = regexp.MustCompile(`\\[A-Za-z]`)

func stripExamples(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, vv := range t {
			if k == "example" {
				continue
			}
			// Drop enums which contain backslash-letter sequences (e.g., \Sent) to avoid generator escaping bugs
			if k == "enum" {
				if arr, ok := vv.([]any); ok {
					stripEnum := false
					for _, ev := range arr {
						if s, ok := ev.(string); ok && backslashLetter.MatchString(s) {
							stripEnum = true
							break
						}
					}
					if stripEnum {
						continue
					}
				}
			}
			out[k] = stripExamples(vv)
		}
		return out
	case []any:
		for i := range t {
			t[i] = stripExamples(t[i])
		}
		return t
	default:
		return v
	}
}


