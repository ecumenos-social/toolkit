package validators

import (
	"context"
	"fmt"
	"regexp"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
)

var (
	emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRegex        = regexp.MustCompile(emailRegexPattern)
)

var _ Validator = ValidateEmail

func ValidateEmail(ctx context.Context, input interface{}, opts ...Option[string]) error {
	v, isStr := input.(string)
	if !isStr {
		return errorwrapper.New("email address must be string")
	}
	if emailRegex.MatchString(v) {
		return nil
	}

	return errorwrapper.New(fmt.Sprintf("invalid email (email: %v)", v))
}
