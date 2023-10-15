package leanauthz

import "strings"

// Effect is the result of an authorization request.
// It can be either ALLOW or DENY.
type Effect string

const (
	// ALLOW means the request is allowed.
	ALLOW Effect = "allow"
	// DENY means the request is denied.
	DENY Effect = "deny"
)

// ParseEffect returns an Effect from a string.
// If the input string is invalid, it returns DENY.
func ParseEffect(str string) Effect {
	switch strings.ToLower(str) {
	case "allow":
		return ALLOW
	case "deny":
		return DENY
	default:
		return DENY
	}
}
