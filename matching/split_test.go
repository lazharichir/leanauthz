package matching_test

import (
	"reflect"
	"testing"

	"github.com/lazharichir/leanauthz/matching"
)

func TestSplitIgnoringEscaped(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      rune
		expected []string
	}{
		{
			name:     "a*b",
			input:    "a\\*b",
			sep:      '*',
			expected: []string{"a\\*b"},
		},
		{
			name:     "no escape characters",
			input:    "a,b,c",
			sep:      ',',
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "escaped separator",
			input:    "a\\,b,c",
			sep:      ',',
			expected: []string{"a\\,b", "c"},
		},
		{
			name:     "escaped escape character",
			input:    "a\\\\,b,c",
			sep:      ',',
			expected: []string{"a\\\\", "b", "c"},
		},
		{
			name:     "multiple escaped separator",
			input:    "a\\,b\\,c,d",
			sep:      ',',
			expected: []string{"a\\,b\\,c", "d"},
		},
		{
			name:     "escaped separator at end",
			input:    "a,b\\,",
			sep:      ',',
			expected: []string{"a", "b\\,"},
		},
		{
			name:     "escaped separator at beginning",
			input:    "\\,a,b",
			sep:      ',',
			expected: []string{"\\,a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := matching.SplitIgnoringEscaped(tt.input, string(tt.sep))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("splitIgnoringEscaped(%q, %q) = %q, expected %q", tt.input, tt.sep, actual, tt.expected)
			}
		})
	}
}
