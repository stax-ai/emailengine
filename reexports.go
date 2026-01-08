package emailengine

import (
	sub "github.com/stax-ai/emailengine/emailengine"
)

// Re-export the SDK types and constructors at the module root for convenient imports:
// import "github.com/stax-ai/emailengine"

type (
	Client      = sub.Client
	Config      = sub.Config
	RetryConfig = sub.RetryConfig
	Option      = sub.Option
)

var (
	New            = sub.New
	WithBaseURL    = sub.WithBaseURL
	WithToken      = sub.WithToken
	WithHTTPClient = sub.WithHTTPClient
)
