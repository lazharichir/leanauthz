package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestNewResource(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		wantErr bool
	}{
		{
			name:    "valid resource",
			str:     "documents/doc_abcdef",
			wantErr: false,
		},
		{
			name:    "empty resource",
			str:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("NewResource() panicked unexpectedly: %v", r)
				}
			}()

			got := authz.NewResource(tt.str)

			if tt.wantErr {
				t.Errorf("NewResource() did not return an error for an empty string")
			}

			if string(got) != tt.str {
				t.Errorf("NewResource() = %v, want %v", got, tt.str)
			}
		})
	}
}

func TestResource_SplitBy(t *testing.T) {
	tests := []struct {
		name string
		r    authz.Resource
		sep  string
		want []string
	}{
		{
			name: "valid resource",
			r:    authz.NewResource("documents/doc_abcdef"),
			sep:  "/",
			want: []string{"documents", "doc_abcdef"},
		},
		{
			name: "multiple pairs of parts",
			r:    authz.NewResource("documents/doc_abcdef/comments/cmt_123456"),
			sep:  "/",
			want: []string{"documents", "doc_abcdef", "comments", "cmt_123456"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.SplitBy(tt.sep)

			if len(got) != len(tt.want) {
				t.Errorf("Resource.SplitBy() returned %d parts, want %d", len(got), len(tt.want))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Resource.SplitBy() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestNewNamespacedResource(t *testing.T) {
	tests := []struct {
		name    string
		ns      string
		id      string
		want    authz.Resource
		wantErr bool
	}{
		{
			name:    "valid resource",
			ns:      "documents",
			id:      "doc_abcdef",
			want:    authz.NewResource("documents/doc_abcdef"),
			wantErr: false,
		},
		{
			name:    "empty namespace",
			ns:      "",
			id:      "doc_abcdef",
			want:    authz.Resource(""),
			wantErr: true,
		},
		{
			name:    "empty id",
			ns:      "documents",
			id:      "",
			want:    authz.Resource(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("NewNamespacedResource() panicked unexpectedly: %v", r)
				}
			}()

			got := authz.NewNamespacedResource(tt.ns, tt.id)

			if tt.wantErr {
				t.Errorf("NewNamespacedResource() did not return an error for an empty string")
			}

			if string(got) != tt.want.String() {
				t.Errorf("NewNamespacedResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
