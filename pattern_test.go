package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestNewPattern(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{
			name: "valid pattern",
			str:  "users:*:read",
		},
		{
			name: "empty pattern",
			str:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authz.NewPattern(tt.str)

			if string(got) != tt.str {
				t.Errorf("NewPattern() = %v, want %v", got, tt.str)
			}
		})
	}
}

func TestPattern_String(t *testing.T) {
	tests := []struct {
		name string
		p    authz.Pattern
		want string
	}{
		{
			name: "valid pattern",
			p:    authz.NewPattern("users:*:read"),
			want: "users:*:read",
		},
		{
			name: "empty pattern",
			p:    authz.NewPattern(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.String()

			if got != tt.want {
				t.Errorf("Pattern.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
