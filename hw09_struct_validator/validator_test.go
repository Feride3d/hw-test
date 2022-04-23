package hw09structvalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		//	meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidateErrors(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr ValidationErrors
	}{
		{in: App{Version: "42.1"}, expectedErr: ValidationErrors{{Field: "Version", Err: ErrInvalidLen}}},
		{in: App{Version: "42.2.2"}, expectedErr: ValidationErrors{{Field: "Version", Err: ErrInvalidLen}}},
		{in: User{Email: "tuftuftuf"}, expectedErr: ValidationErrors{{Field: "Email", Err: ErrRegexNotMatch}}},
		{in: User{ID: "200200"}, expectedErr: ValidationErrors{{Field: "ID", Err: ErrInvalidLen}}},
		{in: User{Age: 80}, expectedErr: ValidationErrors{{Field: "Age", Err: ErrMaxValue}}},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.in)
			require.Error(t, err)

			// unpack errors to return them in ValidationErrors type
			var vErr ValidationErrors
			require.ErrorAs(t, err, &vErr)

			require.EqualError(t, vErr, tt.expectedErr.Error())
		})
	}
}

func TestPositiveValidation(t *testing.T) {
	tests := []interface{}{
		App{Version: "2.2.1"},
		User{Email: "Kate@gmail.com"},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			require.NoError(t, Validate(tt))
		})
	}
}
