package util

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Capitalize the first letter in a string.
// Will ensure all usernames start with a capital letter.
func CapitalizeFirst(text string) (string, error) {
	// Remove whitespace from input string.
	text = strings.TrimSpace(text)

	if text == "" {
		return "", errors.New("empty username")
	}

	// Return input string if first letter is a multi-byte rune
	if _, size := utf8.DecodeRuneInString(text); size > 1 {
		return text, errors.New("multi-byte rune at first letter")
	}

	return strings.ToUpper(string(text[0])) + text[1:], nil
}
