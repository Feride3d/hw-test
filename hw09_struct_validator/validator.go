package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

var valErr ValidationErrors

func (v ValidationErrors) Error() string {
	var strBuild strings.Builder

	for _, validationError := range v {
		strBuild.WriteString(validationError.Error())
	}
	return strBuild.String()
}

// Check that the input interface is a struct
func Validate(v interface{}) error {
	refValue := reflect.ValueOf(v)
	//	refType := refValue.Type()

	if refValue.Kind() != reflect.Struct {
		return errors.New("error: input interface is not a struct")
	}
	return nil
}

func NewError(field string, err error) *ValidationError {
	return &ValidationError{
		Field: field,
		Err:   err,
	}
	return nil
}
