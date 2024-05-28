package validators

import "strings"

type Option[T interface{}] func(...T) []T

var IgnoreCaseOption Option[string] = func(values ...string) []string {
	if len(values) == 0 {
		return values
	}

	out := make([]string, 0, len(values))
	for _, val := range values {
		out = append(out, strings.ToLower(val))
	}

	return out
}
