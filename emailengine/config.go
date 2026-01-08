package emailengine

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Config struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	Logger     *zap.Logger
	Tracer     trace.Tracer
	Retry      RetryConfig
	Timeout    time.Duration
}

type RetryConfig struct {
	MaxAttempts int
	BaseDelayMs int
	MaxDelayMs  int
}

func defaultConfig() Config {
	return Config{
		BaseURL: "https://emailengine.stax.ai",
		Retry: RetryConfig{
			MaxAttempts: 5,
			BaseDelayMs: 200,
			MaxDelayMs:  3000,
		},
	}
}
