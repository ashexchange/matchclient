package spotapi

import "testing"

var mc = NewMarketClient(cli)

func TestMarketList(t *testing.T) {
	mc.MarketList(ctx)
}
