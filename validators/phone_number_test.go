package validators_test

import (
	"context"
	"testing"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/toolkit/primitives"
	"github.com/ecumenos-social/toolkit/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidatePhoneNumber(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		input  interface{}
		opts   []validators.Option[string]
		output error
	}{
		"should works for valid for +x xxx-xxx-xxxx": {
			input:  "+1 234-567-8900",
			output: nil,
		},
		"should works for valid for (xxx) xxx-xxxx": {
			input:  "(123) 456-7890",
			output: nil,
		},
		"should works for valid for xxxxxxxxxx": {
			input:  "1234567890",
			output: nil,
		},
		"should works for valid for +xxxxxxxxxx": {
			input:  "+380504327542",
			output: nil,
		},
		"should works for valid for xxx-xxx-xxxx": {
			input:  "123-456-7890",
			output: nil,
		},
		"should works for valid for +xx xx xxxx xxxx": {
			input:  "+44 20 7123 1234",
			output: nil,
		},
		"should not works for invalid": {
			input:  "example",
			output: errorwrapper.New("invalid phone number (phone number: example)"),
		},
		"should not works for not string": {
			input:  1,
			output: errorwrapper.New("phone number must be string"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := validators.ValidatePhoneNumber(ctx, test.input, test.opts...)
			assert.Truef(t, primitives.IsSameError(test.output, result), "err: %w", result)
		})
	}
}
