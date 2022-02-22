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
	isRune := false

	// Search by runes in a string. IsDigit reports whether the rune is a decimal digit.
	// If an invalid string was passed, the function returns "invalid string".
	for _, r := range packedString {
		switch {
		case unicode.IsDigit(r):
			if !isRune {
				return "", ErrInvalidString
			}
			for i := rune(0); i < r-'0'; i++ {
				unpacked.WriteRune(actualRune)
			}
			isRune = false
		default:
			if isRune {
				unpacked.WriteRune(actualRune)
			}
			actualRune = r
			isRune = true
		}
	}
	if isRune := false; isRune {
		unpacked.WriteRune(actualRune)
	}

	return unpacked.String(), nil
}
