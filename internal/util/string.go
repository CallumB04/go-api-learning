package util

import (
	"strings"
)

// Capitalize the first letter in a string.
// Will ensure all usernames start with a capital letter.
func CapitalizeFirst(text string) string {

	if text == "" {
		return ""
	}

	// Remove whitespace from input string.
	textTrimmed := strings.TrimSpace(text)

	return strings.ToUpper(string(textTrimmed[0])) + textTrimmed[1:]
}
