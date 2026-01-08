package util

import "testing"

func TestRedactToken(t *testing.T) {
	got := RedactToken("abcdefghijklmnop")
	if got != "abcd********mnop" {
		t.Fatalf("unexpected redaction: %s", got)
	}
	if RedactToken("short") != "****" {
		t.Fatalf("expected **** for short token")
	}
}
