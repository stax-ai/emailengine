package emailengine

import "net/http"

type Option func(*Config)

func WithBaseURL(u string) Option {
	return func(c *Config) { c.BaseURL = u }
}

func WithToken(token string) Option {
	return func(c *Config) { c.Token = token }
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Config) { c.HTTPClient = hc }
}
