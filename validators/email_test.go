package validators_test

import (
	"context"
	"testing"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/toolkit/primitives"
	"github.com/ecumenos-social/toolkit/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		input  interface{}
		opts   []validators.Option[string]
		output error
	}{
		"should works for valid": {
			input:  "example@mail.com",
			output: nil,
		},
		"should not works for invalid": {
			input:  "example",
			output: errorwrapper.New("invalid email (email: example)"),
		},
		"should not works for not string": {
			input:  1,
			output: errorwrapper.New("email address must be string"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := validators.ValidateEmail(ctx, test.input, test.opts...)
			assert.Truef(t, primitives.IsSameError(test.output, result), "err: %w", result)
		})
	}
}
