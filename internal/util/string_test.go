package util

import (
	"testing"
)

func TestCapitalizeFirst(t *testing.T) {
	// Create test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"single lower", "a", "A"},
		{"single upper", "A", "A"},
		{"word lower", "hello", "Hello"},
		{"already capped", "Hello", "Hello"},
		{"punctuation first", "!bang", "!bang"},
		{"digit first", "1abc", "1abc"},
		{"leading space", " hello", "Hello"},
		{"multibyte first", "école", "École"},
	}

	// Iterate through unit tests
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if result := CapitalizeFirst(testCase.input); result != testCase.expected {
				t.Errorf("CapitalizeFirst(%q) = %q; want %q", testCase.input, result, testCase.expected)
			}
		})
	}
}
