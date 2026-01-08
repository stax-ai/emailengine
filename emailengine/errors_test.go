package emailengine

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestDecodeError_ParsesJSONFields(t *testing.T) {
	body := []byte(`{"error":"rate_limited","message":"Slow down"}`)
	res := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Body:       io.NopCloser(bytes.NewReader(body)),
		Header:     make(http.Header),
	}
	err := decodeError("op", res)
	ee, ok := err.(*EEError)
	if !ok {
		t.Fatalf("expected EEError, got %T", err)
	}
	if ee.Status != http.StatusTooManyRequests {
		t.Fatalf("status mismatch: %d", ee.Status)
	}
	if ee.Code != "rate_limited" {
		t.Fatalf("code mismatch: %s", ee.Code)
	}
	if ee.Message != "Slow down" {
		t.Fatalf("message mismatch: %s", ee.Message)
	}
	if !ee.Retryable {
		t.Fatalf("expected retryable")
	}
}
