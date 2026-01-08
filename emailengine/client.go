package emailengine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stax-ai/emailengine/api"
	"github.com/stax-ai/emailengine/emailengine/internal/httptransport"
	"github.com/stax-ai/emailengine/emailengine/middleware"
	"github.com/stax-ai/emailengine/emailengine/retry"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Client struct {
	cfg        Config
	httpClient *http.Client
	api        *api.APIClient
	logger     *zap.Logger
	tracer     trace.Tracer
}

func New(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	for _, o := range opts {
		o(&cfg)
	}
	if cfg.Token == "" {
		return nil, fmt.Errorf("token is required")
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	// Build transport with middleware
	ms := []httptransport.RoundTripperMiddleware{
		httptransport.WithRequestTimeoutHeader(),
	}
	// logging
	var logger *zap.Logger
	if cfg.Logger != nil {
		logger = cfg.Logger
	} else {
		logger = zap.NewNop()
	}
	ms = append(ms, middleware.ZapLogging(logger))
	// tracing
	ms = append(ms, middleware.OTelTracing(cfg.Tracer))
	// retry
	rp := retry.Default()
	if cfg.Retry.MaxAttempts > 0 {
		rp.MaxAttempts = cfg.Retry.MaxAttempts
	}
	if cfg.Retry.BaseDelayMs > 0 {
		rp.BaseDelay = time.Duration(cfg.Retry.BaseDelayMs) * time.Millisecond
	}
	if cfg.Retry.MaxDelayMs > 0 {
		rp.MaxDelay = time.Duration(cfg.Retry.MaxDelayMs) * time.Millisecond
	}
	ms = append(ms, retry.Middleware(rp))

	cfg.HTTPClient.Transport = httptransport.New(cfg.HTTPClient.Transport, ms...)

	// Configure generated client
	conf := api.NewConfiguration()
	conf.Servers = api.ServerConfigurations{
		{URL: cfg.BaseURL},
	}
	conf.HTTPClient = cfg.HTTPClient
	// auth via bearer
	conf.AddDefaultHeader("Authorization", "Bearer "+cfg.Token)

	apiClient := api.NewAPIClient(conf)

	return &Client{
		cfg:        cfg,
		httpClient: cfg.HTTPClient,
		api:        apiClient,
		logger:     logger,
		tracer:     cfg.Tracer,
	}, nil
}

// Expose underlying generated clients when needed
func (c *Client) Raw() *api.APIClient { return c.api }
