package types

import (
	"time"
)

type RateLimit struct {
	MaxRequests int64
	Interval    time.Duration
}
