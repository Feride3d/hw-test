package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack function performs a primitive unpacking of a string containing repeated characters.
func Unpack(packedString string) (string, error) {
	// If a packedString has no value, return empty string ""
	if packedString == "" {
		return "", nil
	}

	// A new variables with scope in a function.
	unpacked := strings.Builder{}
	var actualRune rune
	isrune := false

	// Search by runes in a string. IsDigit reports whether the rune is a decimal digit.
	// If an invalid string was passed, the function returns "invalid string".
	for _, r := range packedString {
		switch {
		case unicode.IsDigit(r):
			if !isrune {
				return "", ErrInvalidString
			}
			for i := rune(0); i < r-'0'; i++ {
				unpacked.WriteRune(actualRune)
			}
			isrune = false
		default:
			if isrune {
				unpacked.WriteRune(actualRune)
			}
			actualRune = r
			isrune = true
		}
	}
	if isrune {
		unpacked.WriteRune(actualRune)
	}

	return unpacked.String(), nil
}
