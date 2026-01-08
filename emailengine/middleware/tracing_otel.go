package middleware

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func OTelTracing(tracer trace.Tracer) Middleware {
	if tracer == nil {
		tracer = otel.Tracer("emailengine/sdk")
	}
	return func(next http.RoundTripper) http.RoundTripper {
		return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			ctx, span := tracer.Start(req.Context(), "http.request", trace.WithSpanKind(trace.SpanKindClient))
			defer span.End()
			req = req.WithContext(ctx)
			start := time.Now()
			res, err := next.RoundTrip(req)
			span.SetAttributes(
				attribute.String("http.method", req.Method),
				attribute.String("http.url", req.URL.String()),
			)
			if res != nil {
				span.SetAttributes(attribute.Int("http.status_code", res.StatusCode))
			}
			span.SetAttributes(attribute.Int64("http.duration_ms", time.Since(start).Milliseconds()))
			if err != nil {
				span.RecordError(err)
			}
			return res, err
		})
	}
}
