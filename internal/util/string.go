package util

import (
	"strings"
	"unicode/utf8"
)

// Capitalize the first letter in a string.
// Will ensure all usernames start with a capital letter.
func CapitalizeFirst(text string) string {
	// Remove whitespace from input string.
	text = strings.TrimSpace(text)

	if text == "" {
		return ""
	}

	// Return input string if first letter is a multi-byte rune
	if _, size := utf8.DecodeRuneInString(text); size > 1 {
		return text
	}

	return strings.ToUpper(string(text[0])) + text[1:]
}
