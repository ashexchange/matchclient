package futuresapi

import (
	"context"
)

type ConfigClient struct {
	invoker Invoker
}

func NewConfigClient(invoker Invoker) ConfigClient {
	return ConfigClient{invoker}
}

func (c ConfigClient) UpdateAsset(ctx context.Context) error {
	return c.invoker.Invoke(ctx, "config.update_asset", nil)
}

func (c ConfigClient) UpdateMarket(ctx context.Context) error {
	return c.invoker.Invoke(ctx, "config.update_market", nil)
}

func (c ConfigClient) UpdateIndex(ctx context.Context) error {
	return c.invoker.Invoke(ctx, "config.update_index", nil)
}
