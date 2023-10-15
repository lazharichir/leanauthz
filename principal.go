package leanauthz

import "strings"

// Principal is a string representing the identity of a user, service, or role.
// It is usually in the form of <type>/<id> but this is up to your application.
// e.g., "users/alice", "services/s3.amazonaws.com", "roles/admin", "accounts/123456789012/teams/sales"
type Principal string

// NewPrincipal returns a new Principal. It cannot be empty.
func NewPrincipal(str string) Principal {
	if str == "" {
		panic("authz: principal string cannot be empty")
	}
	return Principal(str)
}

// NewNamespacesPrincipal returns a new Principal with the given namespace.
func NewNamespacedPrincipal(ns, id string) Principal {
	if ns == "" {
		panic("authz: principal namespace string cannot be empty")
	}
	if id == "" {
		panic("authz: principal id string cannot be empty")
	}
	return Principal(strings.Join([]string{ns, id}, `/`))
}

// NewUserPrincipal returns a new Principal with the "users" namespace.
// e.g., "users/alice", "users/bob", "users/123456789012"
func NewUserPrincipal(userID string) Principal {
	return NewNamespacedPrincipal(`users`, userID)
}

// NewServicePrincipal returns a new Principal with the "services" namespace.
// e.g., "services/autocomplete", "services/iam"
func NewServicePrincipal(serviceName string) Principal {
	return NewNamespacedPrincipal(`services`, serviceName)
}

// NewRolePrincipal returns a new Principal with the "roles" namespace.
// e.g., "roles/admin", "roles/manager", "roles/rol_abcdef"
func NewRolePrincipal(roleID string) Principal {
	return NewNamespacedPrincipal(`roles`, roleID)
}
