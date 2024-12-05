package spotapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2"
	"github.com/ashexchange/matchclient/v2/types"
)

type ConfigClient interface {
	// ConfigUpdateAsset 更新币种配置
	ConfigUpdateAsset(ctx context.Context) (*types.Result, error)

	// ConfigUpdateMarket 更新市场配置
	ConfigUpdateMarket(ctx context.Context) (*types.Result, error)

	// ConfigUpdateIndex 更新指数配置
	ConfigUpdateIndex(ctx context.Context) (*types.Result, error)
}

type configClient struct {
	cli matchclient.Invoker
}

var _ ConfigClient = (*configClient)(nil)

func NewConfigClient(c matchclient.Invoker) ConfigClient {
	return &configClient{cli: c}
}

func (c configClient) ConfigUpdateAsset(ctx context.Context) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "config.update_asset", &reply)
	return
}

func (c configClient) ConfigUpdateMarket(ctx context.Context) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "config.update_market", &reply)
	return
}

func (c configClient) ConfigUpdateIndex(ctx context.Context) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "config.update_index", &reply)
	return
}
