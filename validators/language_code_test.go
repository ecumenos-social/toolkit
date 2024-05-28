package validators_test

import (
	"context"
	"testing"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/toolkit/primitives"
	"github.com/ecumenos-social/toolkit/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidateLanguageCode(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		input  interface{}
		opts   []validators.Option[string]
		output error
	}{
		"should works for valid without options": {
			input:  "UKR",
			output: nil,
		},
		"should works for valid with ignore case option": {
			input:  "ukr",
			opts:   []validators.Option[string]{validators.IgnoreCaseOption},
			output: nil,
		},
		"should not works for invalid but with 3 length": {
			input:  "YYY",
			output: errorwrapper.New("unknown language code (code: YYY)"),
		},
		"should not works for invalid but with no-3 length": {
			input:  "YYYY",
			output: errorwrapper.New("language code must contain 3 symbols"),
		},
		"should not works for not string": {
			input:  1,
			output: errorwrapper.New("language code must be string"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := validators.ValidateLanguageCode(ctx, test.input, test.opts...)
			assert.Truef(t, primitives.IsSameError(test.output, result), "err: %w", result)
		})
	}
}
