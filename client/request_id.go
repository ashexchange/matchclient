package client

import (
	"context"
)

type requestId struct{}

func WithRequestId(ctx context.Context, id uint32) context.Context {
	return context.WithValue(ctx, requestId{}, id)
}

func FromContext(ctx context.Context) uint32 {
	id, _ := ctx.Value(requestId{}).(uint32)
	return id
}
