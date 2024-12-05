package futuresapi

import (
	"context"
	"testing"

	"github.com/ashexchange/matchclient/v2/client"
)

const (
	endpoint = "http://accesshttp.aie.test"
	userId1  = 1000
	userId2  = 1001
)

var (
	ctx = context.Background()
	cli = client.NewClient(endpoint)
)

func call0[RESP any](t *testing.T, fn func(ctx context.Context) (RESP, error)) RESP {
	var (
		resp RESP
		err  error
	)
	resp, err = fn(ctx)
	if err != nil {
		t.Errorf("invoke error: %s", err)
		return resp
	}

	t.Logf("%+v", resp)

	return resp
}

func call1[REQ any](t *testing.T, fn func(ctx context.Context, req REQ) error, req REQ) {
	if err := fn(ctx, req); err != nil {
		t.Errorf("invoke error: %s", err)
	}
}

func call2[REQ, RESP any](t *testing.T, fn func(ctx context.Context, req REQ) (RESP, error), req REQ) RESP {
	var (
		resp RESP
		err  error
	)
	resp, err = fn(ctx, req)
	if err != nil {
		t.Errorf("invoke error: %s", err)
		return resp
	}

	t.Logf("%+v", resp)

	return resp
}

func foreach[T any](t *testing.T, list []T) {
	for _, item := range list {
		t.Logf("%+v", item)
	}
}
