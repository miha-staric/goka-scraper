// Package utils holds utilities/helpers
package utils

import "strings"

// NormalizeSpaces collapses consecutive spaces into one.
// Leading/trailing spaces are preserved.
func NormalizeSpaces(s string) string {
	var b strings.Builder
	space := false
	for _, r := range s {
		if r == ' ' {
			if !space {
				b.WriteRune(' ')
				space = true
			}
		} else {
			b.WriteRune(r)
			space = false
		}
	}
	return b.String()
}
