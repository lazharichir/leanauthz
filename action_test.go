package leanauthz_test

import (
	"testing"

	"github.com/lazharichir/leanauthz"
)

func TestNewAction(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		wantErr bool
	}{
		{
			name:    "valid action",
			str:     "read",
			wantErr: false,
		},
		{
			name:    "empty action",
			str:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("NewAction() panicked unexpectedly: %v", r)
				}
			}()

			got := leanauthz.NewAction(tt.str)

			if tt.wantErr {
				t.Errorf("NewAction() did not return an error for an empty string")
			}

			if string(got) != tt.str {
				t.Errorf("NewAction() = %v, want %v", got, tt.str)
			}
		})
	}
}
