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

	return strings.ToUpper(string(text[0])) + text[1:]
}
