package environment

import (
	"context"
)

type Process interface {
	Running(ctx context.Context) (bool, error)
	Exists(ctx context.Context) (bool, error)
	Configure(ctx context.Context) error
}
