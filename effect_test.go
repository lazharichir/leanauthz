package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestParseEffect(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want authz.Effect
	}{
		{
			name: "valid allow",
			str:  "allow",
			want: authz.ALLOW,
		},
		{
			name: "valid deny",
			str:  "deny",
			want: authz.DENY,
		},
		{
			name: "valid allow in caps",
			str:  "ALLOW",
			want: authz.ALLOW,
		},
		{
			name: "valid deny in caps",
			str:  "DENY",
			want: authz.DENY,
		},
		{
			name: "invalid effect",
			str:  "invalid",
			want: authz.DENY,
		},
		{
			name: "invalid effect (whitespace)",
			str:  "allow ",
			want: authz.DENY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authz.ParseEffect(tt.str)

			if got != tt.want {
				t.Errorf("ParseEffect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEffect_Allow(t *testing.T) {
	if authz.ALLOW != "allow" {
		t.Errorf("ALLOW constant is not equal to 'allow'")
	}
}

func TestEffect_Deny(t *testing.T) {
	if authz.DENY != "deny" {
		t.Errorf("DENY constant is not equal to 'deny'")
	}
}
