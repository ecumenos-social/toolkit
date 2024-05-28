package validators

import (
	"context"
	"fmt"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/schemas/constants"
)

var _ Validator = ValidateCountryCode

func ValidateCountryCode(ctx context.Context, input interface{}, opts ...Option[string]) error {
	v, isStr := input.(string)
	if !isStr {
		return errorwrapper.New("country code must be string")
	}
	if len(v) != 3 {
		return errorwrapper.New("country code must contain 3 symbols")
	}

	countries, err := constants.LoadCountries(ctx)
	if err != nil {
		return errorwrapper.WrapMessage(err, "failed load countries")
	}

	for _, cc := range countries {
		val1, val2 := string(cc.Alpha3), v
		for _, opt := range opts {
			values := opt(val1, val2)
			val1 = values[0]
			val2 = values[1]
		}
		if val1 == val2 {
			return nil
		}
	}

	return errorwrapper.New(fmt.Sprintf("unknown country code (code: %v)", v))
}
