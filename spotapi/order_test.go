package spotapi

import (
	"testing"

	"github.com/ashexchange/matchclient/v2/types"
)

var oc = NewOrderClient(cli)

func TestOrderClient_Depth(t *testing.T) {
	asset := "USDT"
	oc.OrderPutLimit(ctx, &OrderPutLimitRequest{
		UserId:       1000,
		Account:      0,
		Market:       "BTC-USDT",
		Side:         types.Buy,
		Amount:       "0.01",
		Price:        "65000",
		FeeAsset:     &asset,
		TakerFeeRate: "0.0002",
		MakerFeeRate: "0.0001",
		Source:       "TEST",
		FeeDiscount:  "1",
	})

	oc.OrderDepth(ctx, &OrderDepthRequest{
		Market:   "BTC-USDT",
		Limit:    1,
		Interval: "60",
	})
}
