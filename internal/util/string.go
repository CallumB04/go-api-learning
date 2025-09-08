package util

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrEmptyUsername    = errors.New("empty username")
	ErrMultibyteFirst   = errors.New("multi-byte rune at first character")
	ErrInvalidUTF8First = errors.New("invalid utf-8 at first byte")
)

// Capitalize the first letter in a string.
// Will ensure all usernames start with a capital letter.
func CapitalizeFirst(text string) (string, error) {
	// Remove whitespace from input string.
	text = strings.TrimSpace(text)

	if text == "" {
		return "", ErrEmptyUsername
	}

	// Decode first letter of input string into rune
	firstRune, size := utf8.DecodeRuneInString(text)

	// Invalid UTF-8 at the first byte of the string
	if firstRune == utf8.RuneError && size == 1 {
		return text, ErrInvalidUTF8First
	}

	// Multi-byte rune at the first character of the string
	if size > 1 {
		return text, ErrMultibyteFirst
	}

	return strings.ToUpper(string(text[0])) + text[1:], nil
}
