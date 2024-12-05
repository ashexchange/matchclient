package futuresapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2/types"
)

type IndexClient struct {
	invoker Invoker
}

func NewIndexClient(invoker Invoker) IndexClient {
	return IndexClient{invoker}
}

type IndexDebugRequest struct {
	Market     string
	Debug      bool
	IndexPrice string
	SignPrice  *string
}

func (c IndexClient) Debug(ctx context.Context, req IndexDebugRequest) error {
	params := []any{
		req.Market,
		req.Debug,
		req.IndexPrice,
	}
	if req.SignPrice != nil {
		params = append(params, req.SignPrice)
	}

	return c.invoker.Invoke(ctx, "index.debug", nil, params...)
}

type IndexDetail struct {
	Index   types.Number   `json:"index"`
	Time    types.Time     `json:"time"`
	Sources []*IndexSource `json:"sources"`
}

type IndexSource struct {
	Exchange string     `json:"exchange"`
	Weight   string     `json:"weight"`
	Time     types.Time `json:"time"`
}

type IndexListResult map[string]*IndexDetail

func (c IndexClient) List(ctx context.Context) (result IndexListResult, err error) {
	err = c.invoker.Invoke(ctx, "index.list", &result)

	return
}

type IndexQueryRequest struct {
	Market string
}

type IndexInfoResult struct {
	Index types.Number `json:"index"`
	Name  string       `json:"name"`
	Time  types.Time   `json:"time"`
}

func (c IndexClient) Query(ctx context.Context, req IndexQueryRequest) (result *IndexInfoResult, err error) {
	err = c.invoker.Invoke(ctx, "index.query", &result, req.Market)

	return
}
