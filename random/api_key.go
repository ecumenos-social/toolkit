package random

import (
	"strings"

	"github.com/ecumenos-social/toolkit/hash"
)

func GenAPIKey(scope, node string) (string, error) {
	h, err := hash.MD5(GenNumericString(100))
	if err != nil {
		return "", err
	}
	return strings.Join([]string{scope, hash.CRC32(node), h}, "."), nil
}
