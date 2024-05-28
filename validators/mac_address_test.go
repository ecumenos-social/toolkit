package validators_test

import (
	"context"
	"testing"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/toolkit/primitives"
	"github.com/ecumenos-social/toolkit/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidateMACAddress(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		input  interface{}
		opts   []validators.Option[string]
		output error
	}{
		"should works for valid xx:xx:xx:xx:xx:xx format": {
			input:  "00:00:5e:00:53:01",
			output: nil,
		},
		"should works for valid xx:xx:xx:xx:xx:xx:xx:xx format": {
			input:  "02:00:5e:10:00:00:00:01",
			output: nil,
		},
		"should works for valid xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx format": {
			input:  "00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01",
			output: nil,
		},
		"should works for valid xx-xx-xx-xx-xx-xx format": {
			input:  "00-00-5e-00-53-01",
			output: nil,
		},
		"should works for valid xx-xx-xx-xx-xx-xx-xx-xx format": {
			input:  "02-00-5e-10-00-00-00-01",
			output: nil,
		},
		"should works for valid xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx-xx format": {
			input:  "00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01",
			output: nil,
		},
		"should works for valid xxxx.xxxx.xxxx format": {
			input:  "0000.5e00.5301",
			output: nil,
		},
		"should works for valid xxxx.xxxx.xxxx.xxxx format": {
			input:  "0200.5e10.0000.0001",
			output: nil,
		},
		"should works for valid xxxx.xxxx.xxxx.xxxx.xxxx.xxxx.xxxx.xxxx.xxxx.xxxx format": {
			input:  "0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001",
			output: nil,
		},
		"should not works for invalid": {
			input:  "xxxxxxxxxx",
			output: errorwrapper.New("invalid mac address (mac address: xxxxxxxxxx)"),
		},
		"should not works for not string": {
			input:  1,
			output: errorwrapper.New("mac address must be string"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := validators.ValidateMACAddress(ctx, test.input, test.opts...)
			assert.Truef(t, primitives.IsSameError(test.output, result), "err: %w", result)
		})
	}
}
