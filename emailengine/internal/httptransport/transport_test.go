package httptransport

import (
	"context"
	"net/http"
	"testing"
	"time"
)

type rtFunc func(*http.Request) (*http.Response, error)

func (f rtFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestWithRequestTimeoutHeader_SetsHeaderFromContextDeadline(t *testing.T) {
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	called := false
	base := rtFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		got := req.Header.Get("X-EE-Timeout")
		if got == "" {
			t.Fatalf("expected X-EE-Timeout to be set")
		}
		return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header)}, nil
	})

	tr := New(base, WithRequestTimeoutHeader())
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.test", nil)
	_, _ = tr.RoundTrip(req)
	if !called {
		t.Fatalf("base RoundTripper not called")
	}
}
