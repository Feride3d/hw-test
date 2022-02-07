package hw02unpackstring

import (
	"errors"
	"strconv"
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
	// Convert the packedString into a slice of runes.
	sliceRunes := []rune(packedString)
	// A new variable of type strings.Builder
	var unpacked strings.Builder
	// Search by the slice of runes.
	for i := range sliceRunes {
		actualRune := sliceRunes[i]
		nextRune := sliceRunes[i+1]
		switch {
		// If an invalid string was passed (the first rune is a digit), the function returns an error.
		case (i == 0 && unicode.IsDigit(actualRune)): // IsDigit reports whether the rune is a decimal digit.
			return "", ErrInvalidString
		/* If an invalid string was passed (the actual rune is a digit and the next rune
		(a rune following the actual rune) is a digit), the function returns an error. */
		case (unicode.IsDigit(actualRune) && unicode.IsDigit(nextRune)):
			return "", ErrInvalidString
		}
		/* If the the next rune (a rune following the actual rune) is a digit,
		then repeat the actual rune in the amount of the next rune. */
		if unicode.IsLetter(actualRune) && unicode.IsDigit(nextRune) {
			repeatRune, err := strconv.Atoi(string(nextRune))
			if err != nil {
				return "", err
			}
			unpacked.WriteString(strings.Repeat(string(actualRune), repeatRune))
		} else {
			unpacked.WriteRune(actualRune)
		}
	}
	return unpacked.String(), nil
}
