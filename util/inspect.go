package util

import "unicode"

// IsFirstUpper ...
func IsFirstUpper(v string) bool {
	if v == "" {
		return false
	}
	r := rune(v[0])
	return unicode.IsUpper(r)
}
