package hw02unpackstring

import (
	"errors"
	"unicode"
	"strings"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

// Функция Unpack осуществляет примитивную распаковку строки, содержащую повторяющиеся символы
func Unpack(packedString string) (string, error) {
	// если передается строка без значения, взвращаем пустую строку ""
	if packedString == "" {
		return "", nil
	}
	// конвертируем передаваемую строку в слайс рун 
	sliceRunes := []rune(packedString)
	// Вводим переменную типа strings.Builder
	var unpacked strings.Builder 
	// Поиск по слайсу рун
	for i := range sliceRunes {
		actualRune := sliceRunes[i]
		nextRune := sliceRunes[i+1]
		switch {
		// Если была передана некорректная строка (первый символ цифра), функция возвращает ошибку
		case (i == 0 && unicode.IsDigit(actualRune)): // IsDigit reports whether the rune is a decimal digit.
			return "", ErrInvalidString
		// Если была передана некорректная строка (акуальный символ является цифрой и следующий за ним символ яляется цифрой), функция возвращает ошибку
		case (unicode.IsDigit(actualRune) && unicode.IsDigit(nextRune)):
			return "", ErrInvalidString
		// Если символ, следующий за актуальным символом, является цифрой, то повторяем актуальный символ в количестве указанной цифры
		case unicode.IsLetter(actualRune) && unicode.IsDigit(nextRune):
			repeatRune, err := strconv.Atoi(string(nextRune))
			if err != nil {
				return "", err
			}
			unpacked.WriteString(strings.Repeat(string(actualRune), repeatRune))
		} else {
			unpacked.WriteRune(actualRune)
		}
	}
}