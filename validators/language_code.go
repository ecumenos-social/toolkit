package validators

import (
	"context"
	"fmt"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
	"github.com/ecumenos-social/schemas/constants"
)

var _ Validator = ValidateLanguageCode

func ValidateLanguageCode(ctx context.Context, input interface{}, opts ...Option[string]) error {
	v, isStr := input.(string)
	if !isStr {
		return errorwrapper.New("language code must be string")
	}
	if len(v) != 3 {
		return errorwrapper.New("language code must contain 3 symbols")
	}

	languages, err := constants.LoadLanguages(ctx)
	if err != nil {
		return errorwrapper.WrapMessage(err, "failed load languages")
	}

	for _, lang := range languages {
		val1, val2 := string(lang.Alpha3), v
		for _, opt := range opts {
			values := opt(val1, val2)
			val1 = values[0]
			val2 = values[1]
		}
		if val1 == val2 {
			return nil
		}
	}

	return errorwrapper.New(fmt.Sprintf("unknown language code (code: %v)", v))
}
