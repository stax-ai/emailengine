package httptransport

import (
	"net/http"
	"strconv"
	"time"
)

type RoundTripperMiddleware func(next http.RoundTripper) http.RoundTripper

type Transport struct {
	base        http.RoundTripper
	middlewares []RoundTripperMiddleware
}

func New(base http.RoundTripper, middlewares ...RoundTripperMiddleware) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Transport{base: base, middlewares: middlewares}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.base
	for i := len(t.middlewares) - 1; i >= 0; i-- {
		rt = t.middlewares[i](rt)
	}
	return rt.RoundTrip(req)
}

// WithRequestTimeoutHeader maps context deadline to X-EE-Timeout (ms)
func WithRequestTimeoutHeader() RoundTripperMiddleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if deadline, ok := req.Context().Deadline(); ok {
				ms := time.Until(deadline).Milliseconds()
				if ms > 0 {
					// set header directly; server expects milliseconds
					req.Header.Set("X-EE-Timeout", itoa(ms))
				}
			}
			return next.RoundTrip(req)
		})
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func itoa(v int64) string {
	return strconv.FormatInt(v, 10)
}
