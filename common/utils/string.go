package utils

import (
	"bytes"
	"unicode"
)

func CamelToSnake(name string) string {
	var ret bytes.Buffer

	multipleUpper := false
	var lastUpper rune
	var beforeUpper rune

	for _, c := range name {
		// Non-lowercase character after uppercase is considered to be uppercase too.
		isUpper := (unicode.IsUpper(c) || (lastUpper != 0 && !unicode.IsLower(c)))

		if lastUpper != 0 {
			// Output a delimiter if last character was either the first uppercase character
			// in a row, or the last one in a row (e.g. 'S' in "HTTPServer").
			// Do not output a delimiter at the beginning of the name.

			firstInRow := !multipleUpper
			lastInRow := !isUpper

			if ret.Len() > 0 && (firstInRow || lastInRow) && beforeUpper != '_' {
				ret.WriteByte('_')
			}
			ret.WriteRune(unicode.ToLower(lastUpper))
		}

		// Buffer uppercase char, do not output it yet as a delimiter may be required if the
		// next character is lowercase.
		if isUpper {
			multipleUpper = (lastUpper != 0)
			lastUpper = c
			continue
		}

		ret.WriteRune(c)
		lastUpper = 0
		beforeUpper = c
		multipleUpper = false
	}

	if lastUpper != 0 {
		ret.WriteRune(unicode.ToLower(lastUpper))
	}
	return ret.String()
}

func StringSliceContainsValue(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
