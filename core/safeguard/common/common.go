package common

import "context"

type GuardFunc func(ctx context.Context) (bool, error)
