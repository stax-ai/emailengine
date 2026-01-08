package util

import "strings"

func RedactToken(s string) string {
	if s == "" {
		return s
	}
	if len(s) <= 8 {
		return "****"
	}
	return s[:4] + strings.Repeat("*", len(s)-8) + s[len(s)-4:]
}
