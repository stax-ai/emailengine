package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func ZapLogging(logger *zap.Logger) Middleware {
	if logger == nil {
		logger = zap.NewNop()
	}
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			start := time.Now()
			res, err := next.RoundTrip(req)
			duration := time.Since(start)
			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
				zap.Duration("duration_ms", duration),
			}
			if res != nil {
				fields = append(fields, zap.Int("status", res.StatusCode))
			}
			if err != nil {
				logger.Error("http.request", append(fields, zap.Error(err))...)
			} else {
				logger.Info("http.request", fields...)
			}
			return res, err
		})
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
