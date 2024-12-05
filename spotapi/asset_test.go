package spotapi

import (
	"testing"
	"time"
)

var ac = NewAssetClient(cli)

func TestAssetList(t *testing.T) {
	ac.AssetList(ctx)
}

func TestAssetClient_Query(t *testing.T) {
	ac.AssetQuery(ctx, &AssetQueryRequest{})
}

func TestAssetClient_Update(t *testing.T) {
	ac.AssetUpdate(ctx, &AssetUpdateRequest{
		UserId:     1000,
		Account:    0,
		Asset:      "USDT",
		Business:   "Deposit",
		BusinessId: uint64(time.Now().Unix()),
		Change:     "1000",
		Detail:     map[string]interface{}{},
	})
}
