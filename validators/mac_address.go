package validators

import (
	"context"
	"fmt"
	"net"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
)

var _ Validator = ValidateMACAddress

func ValidateMACAddress(ctx context.Context, input interface{}, opts ...Option[string]) error {
	v, isStr := input.(string)
	if !isStr {
		return errorwrapper.New("mac address must be string")
	}

	if _, err := net.ParseMAC(v); err != nil {
		return errorwrapper.WrapMessage(err, fmt.Sprintf("invalid mac address (mac address: %v)", v))
	}
	return nil
}
