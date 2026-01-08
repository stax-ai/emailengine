package api

// ReferrerAllowlist represents allowed HTTP referrer patterns for a token.
// Note: The upstream schema permits boolean false as a special sentinel; this SDK treats it as empty list.
type ReferrerAllowlist []string
