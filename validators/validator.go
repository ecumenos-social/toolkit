package validators

import "context"

type Validator func(ctx context.Context, v interface{}, opts ...Option[string]) error
