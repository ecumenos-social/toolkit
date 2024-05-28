package validators

import (
	"context"
	"fmt"
	"regexp"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
)

var (
	phoneRegexPattern       = `^\+?[0-9\s\-\(\)]+$`
	phoneRegex              = regexp.MustCompile(phoneRegexPattern)
	inlinePhoneRegexPattern = `^(\+|)\d{5,15}$`
	inlinePhoneRegex        = regexp.MustCompile(inlinePhoneRegexPattern)
)

var _ Validator = ValidatePhoneNumber

func ValidatePhoneNumber(ctx context.Context, input interface{}, opts ...Option[string]) error {
	v, isStr := input.(string)
	if !isStr {
		return errorwrapper.New("phone number must be string")
	}
	if phoneRegex.MatchString(v) {
		return nil
	}
	if inlinePhoneRegex.MatchString(v) {
		return nil
	}

	return errorwrapper.New(fmt.Sprintf("invalid phone number (phone number: %v)", v))
}
