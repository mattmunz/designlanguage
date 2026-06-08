package strings

import (
	"unicode"
	"unicode/utf8"
)

func LowercaseFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Decode the first rune safely to support UTF-8 characters
	r, size := utf8.DecodeRuneInString(s)

	// Convert the first rune to lowercase and reconstruct the string
	return string(unicode.ToLower(r)) + s[size:]
}
