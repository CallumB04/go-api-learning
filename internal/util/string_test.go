package util

import (
	"testing"
)

func TestCapitalizeFirst(t *testing.T) {
	// Create test cases.
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"empty space", " ", ""},
		{"single lower", "a", "A"},
		{"single upper", "A", "A"},
		{"word lower", "hello", "Hello"},
		{"already capped", "Hello", "Hello"},
		{"punctuation first", "!hello", "!hello"},
		{"digit first", "1hello", "1hello"},
		{"leading space", " hello", "Hello"},
		{"multibyte first", "école", "école"},
		{"leading space and multibyte", " école", "école"},
	}

	// Iterate through unit tests.
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if result := CapitalizeFirst(testCase.input); result != testCase.expected {
				t.Errorf("CapitalizeFirst(%q) = %q; expected %q", testCase.input, result, testCase.expected)
			}
		})
	}
}
