package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestNewNamespacedPrincipal(t *testing.T) {
	tests := []struct {
		name    string
		ns      string
		id      string
		want    authz.Principal
		wantErr bool
	}{
		{
			name:    "valid principal",
			ns:      "users",
			id:      "user123",
			want:    authz.Principal("users/user123"),
			wantErr: false,
		},
		{
			name:    "empty namespace",
			ns:      "",
			id:      "user123",
			want:    authz.Principal(""),
			wantErr: true,
		},
		{
			name:    "empty id",
			ns:      "users",
			id:      "",
			want:    authz.Principal(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("NewNamespacedPrincipal() panicked unexpectedly: %v", r)
				}
			}()

			got := authz.NewNamespacedPrincipal(tt.ns, tt.id)

			if tt.wantErr {
				t.Errorf("NewNamespacedPrincipal() did not return an error for an empty string")
			}

			if string(got) != string(tt.want) {
				t.Errorf("NewNamespacedPrincipal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPrincipal(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		wantErr bool
	}{
		{
			name:    "valid principal",
			str:     "user123",
			wantErr: false,
		},
		{
			name:    "empty principal",
			str:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("NewPrincipal() panicked unexpectedly: %v", r)
				}
			}()

			got := authz.NewPrincipal(tt.str)

			if tt.wantErr {
				t.Errorf("NewPrincipal() did not return an error for an empty string")
			}

			if string(got) != tt.str {
				t.Errorf("NewPrincipal() = %v, want %v", got, tt.str)
			}
		})
	}
}
