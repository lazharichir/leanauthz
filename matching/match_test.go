package matching_test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/lazharichir/leanauthz/matching"
)

func TestMatchEscaping(t *testing.T) {
	{
		pattern := "a\\*b"
		subject := "a\\*b"
		expected := true
		got := matching.Match(pattern, subject)
		if got != expected {
			t.Errorf("Glob(%q, %q) returned %v, expected %v", pattern, subject, !expected, expected)
		}
	}
}

func TestMatch(t *testing.T) {

	tests := []struct {
		pattern  string
		subject  string
		expected bool
	}{
		//
		{"", "hello", false},
		{"**********", "", true},
		{"*", "ab", true},
		{"-*-", "a-b", false},
		{"*-*", "a-b", true},
		{"**", "ReadStuff", true},
		{"****", "a", true},
		//
		{"he*o", "hello", true},
		{"hello", "hello", true},
		{"?", "a", false},
		//
		{"*hello", "hello", true},
		{"*hello", "hellohello", true},
		{"*hello", "hell", false},
		//
		{"hello*", "hello", true},
		{"hello*", "hellohello", true},
		{"hello*", "hell", false},
		//
		{"users/*", "users/john", true},
		{"accounts/*/users/*", "accounts/123/users/john", true},
		//
		{"Read*", "Read", true},
		{"Read*", "ReadDocuments", true},
		{"Read*", "Readbills", true},
		{"read:*", "read:bills", true},
		//
		{"*Invoices", "ReadInvoices", true},
		{"*Invoices", "Readbills", false},
		{"*Invoices", "ReadInvoices", true},
	}

	for _, test := range tests {
		matched := matching.Match(test.pattern, test.subject)
		if matched != test.expected {
			t.Errorf("Glob(%q, %q) returned %v, expected %v", test.pattern, test.subject, !test.expected, test.expected)
		}
	}
}

func TestEscape(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		delimiter string
		expected  string
	}{
		{
			name:      "no delimiter",
			str:       "hello world",
			delimiter: ",",
			expected:  "hello world",
		},
		{
			name:      "single delimiter",
			str:       "hello,world",
			delimiter: ",",
			expected:  "hello\\,world",
		},
		{
			name:      "multiple delimiters",
			str:       "hello,world,how,are,you",
			delimiter: ",",
			expected:  "hello\\,world\\,how\\,are\\,you",
		},
		{
			name:      "delimiter at beginning",
			str:       ",hello,world",
			delimiter: ",",
			expected:  "\\,hello\\,world",
		},
		{
			name:      "delimiter at end",
			str:       "hello,world,",
			delimiter: ",",
			expected:  "hello\\,world\\,",
		},
		{
			name:      "consecutive delimiters",
			str:       "hello**world****",
			delimiter: "**",
			expected:  "hello\\**world\\**\\**",
		},
		{
			name:      "empty string",
			str:       "",
			delimiter: ",",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := matching.Escape(tt.str, tt.delimiter)
			if actual != tt.expected {
				t.Errorf("Escape(%q, %q) = %q, expected %q", tt.str, tt.delimiter, actual, tt.expected)
			}
		})
	}
}

func TestGlobFuzzy(t *testing.T) {
	for i := 0; i < 1000; i++ {
		pattern := randomString(10)
		subj := randomString(10)

		expected := strings.Contains(subj, pattern)

		if matching.Match(pattern, subj) != expected {
			t.Errorf("Glob(%q, %q) returned %v, expected %v", pattern, subj, !expected, expected)
		}
	}
}

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
