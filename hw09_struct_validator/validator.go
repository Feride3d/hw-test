package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidLen    = errors.New("invalid len")
	ErrRegexNotMatch = errors.New("regex not match")
	ErrMaxValue      = errors.New("the number cannot be greater than the max value")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for _, e := range v {
		fmt.Fprintf(&sb, "[f: %s, e: %v]", e.Field, e.Err)
	}
	return sb.String()
}

// Check that the input interface is a struct.
func Validate(v interface{}) error {
	refValue := reflect.ValueOf(v)
	refType := refValue.Type()

	if refValue.Kind() != reflect.Struct {
		return errors.New("input interface is not a struct")
	}

	var vErrs ValidationErrors

	for i := 0; i < refType.NumField(); i++ {
		fieldType := refType.Field(i)

		tagValue := fieldType.Tag.Get("validate")
		if len(tagValue) == 0 {
			continue
		}

		tagFields := strings.Split(tagValue, ":") // разбиваем поле на два по тегу
		if len(tagFields) != 2 {                  // если полей не два, то возвращаем ошибку
			return errors.New("must be two fields")
		}

		validateFn := tagFields[0] // что проверяем (длину, минимальное значение, прочее)
		validateParam := tagFields[1]
		fieldValue := refValue.Field(i)

		var validationErr error

		switch refValue.Kind() {
		case reflect.String:
			validationErr = validateString(validateFn, validateParam, fieldValue.String())
		case reflect.Int:
			//	validationErr = validateInt(validateFn, validateParam, int(fieldValue.Int()))
		}

		if validationErr != nil {
			vErrs = append(vErrs, ValidationError{Field: fieldType.Name, Err: validationErr})
		}
	}

	if len(vErrs) != 0 {
		return vErrs
	}

	return nil
}

func validateString(fn string, validationParam string, refValue string) error {
	switch fn {
	case "len":
		lenValue, err := strconv.Atoi(validationParam) // check len
		if err != nil {
			return nil
		}

		if len(refValue) != lenValue { // panic avoid
			return ErrInvalidLen
		}

	case "regexp":
		re, err := regexp.Compile(validationParam)
		if err != nil {
			return nil // err
		}

		if !re.MatchString(refValue) {
			return ErrRegexNotMatch
		}
	default:
		return nil
	}

	return nil
}

/*func validateInt(fn string, validationParam string, refValue int) error {
	switch fn {
	case "minValue":


		if len(refValue) != minValue { // panic avoid
			return ErrInvalidLen
		}
	default:
		return nil
	}

	return nil
} */
