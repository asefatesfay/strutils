package strutils

import (
	"strings"
	"unicode"
)

// Returns the string changed with uppercase.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Returns the string changed with lowercase.
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// Returns the string changed with the first letter to uppercase.
func ToFirstUpperCase(s string) string {
	if len(s) < 1 {
		return s
	}

	t := strings.Trim(s, " ")
	t = strings.ToLower(t)
	res := []rune(t)
	res[0] = unicode.ToUpper(res[0])
	return string(res)
}
