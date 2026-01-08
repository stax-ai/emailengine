package api

// AddressAllowlist represents a list of IP addresses or CIDR ranges allowed to use a token.
// Note: The upstream schema permits boolean false as a special sentinel; this SDK treats it as empty list.
type AddressAllowlist []string
