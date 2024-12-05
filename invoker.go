package matchclient

import (
	"context"
)

type Invoker interface {
	Invoke(ctx context.Context, method string, result any, params ...any) error
}
