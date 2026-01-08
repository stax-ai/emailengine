package retry

import (
	"errors"
	"math"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type Policy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func Default() Policy {
	return Policy{MaxAttempts: 5, BaseDelay: 200 * time.Millisecond, MaxDelay: 3 * time.Second}
}

func (p Policy) NextDelay(attempt int) time.Duration {
	backoff := float64(p.BaseDelay) * math.Pow(2, float64(attempt))
	jitter := 0.2 + rand.Float64()*0.8
	d := time.Duration(backoff * jitter)
	if d > p.MaxDelay {
		return p.MaxDelay
	}
	return d
}

func ShouldRetry(res *http.Response, err error) bool {
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return true
		}
		// Dial errors etc.
		var opErr *net.OpError
		return errors.As(err, &opErr)
	}
	if res == nil {
		return true
	}
	if res.StatusCode == http.StatusTooManyRequests {
		return true
	}
	if res.StatusCode == http.StatusBadGateway || res.StatusCode == http.StatusServiceUnavailable || res.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}

func Middleware(policy Policy) func(next http.RoundTripper) http.RoundTripper {
	if policy.MaxAttempts <= 0 {
		policy = Default()
	}
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			var res *http.Response
			var err error
			for attempt := 0; attempt < policy.MaxAttempts; attempt++ {
				res, err = next.RoundTrip(req)
				if !ShouldRetry(res, err) {
					return res, err
				}
				// Respect Retry-After if provided
				if res != nil {
					if ra := res.Header.Get("Retry-After"); ra != "" {
						if secs, parseErr := time.ParseDuration(ra + "s"); parseErr == nil {
							time.Sleep(min(secs, policy.MaxDelay))
							continue
						}
					}
				}
				time.Sleep(policy.NextDelay(attempt))
			}
			return res, err
		})
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
