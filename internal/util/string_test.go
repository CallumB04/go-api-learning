package util

import (
	"testing"
)

func TestCapitalizeFirst(t *testing.T) {
	// Create test cases.
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError error
	}{
		{"empty", "", "", ErrEmptyUsername},
		{"empty space", " ", "", ErrEmptyUsername},
		{"single lower", "a", "A", nil},
		{"single upper", "A", "A", nil},
		{"word lower", "hello", "Hello", nil},
		{"already capped", "Hello", "Hello", nil},
		{"punctuation first", "!hello", "!hello", nil},
		{"digit first", "1hello", "1hello", nil},
		{"leading space", " hello", "Hello", nil},
		{"multibyte first", "école", "école", ErrMultibyteFirst},
		{"leading space and multibyte", " école", "école", ErrMultibyteFirst},
	}

	// Iterate through unit tests.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Pass test input into function
			result, err := CapitalizeFirst(test.input)

			// Fail test error received does not match the expected error
			if err != nil {

				if err.Error() != test.expectedError.Error() {
					t.Errorf("expected error '%s'; got error '%s'", test.expectedError.Error(), err.Error())
				}
			}

			// Fail test result does not match the expected value
			if result != test.expected {
				t.Errorf("CapitalizeFirst(%s) = %s; expected %s", test.input, result, test.expected)
			}

		})
	}
}
