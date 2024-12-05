package futuresapi

import (
	"testing"
	"time"

	"github.com/ashexchange/matchclient/v2/types"
)

var oc = NewOrderClient(cli)

func TestOrderClient_PutLimit(t *testing.T) {
	call2(t, oc.PutLimit, OrderPutLimitRequest{
		UserId:       userId1,
		Market:       "BTCUSDT",
		Side:         Short,
		Amount:       "0.2",
		Price:        "31194",
		TakerFeeRate: "0.001",
		MakerFeeRate: "0.001",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		ClientId:     "test",
	})
}

func TestOrderClient_PutMarket(t *testing.T) {
	call2(t, oc.PutMarket, OrderPutMarketRequest{
		UserId:       userId1 + 1,
		Market:       "BTCUSDT",
		Side:         Short,
		Amount:       "0.1",
		TakerFeeRate: "0.001",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		Option:       0,
		ClientId:     "test",
	})
}

func TestOrderClient_LimitClose(t *testing.T) {
	call2(t, oc.LimitClose, OrderLimitCloseRequest{
		UserId:       100,
		Market:       "BTCUSDT",
		PositionId:   214,
		Price:        "30500",
		TakerFeeRate: "0.001",
		MakerFeeRate: "0.001",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		EffectType:   1,
		Option:       0,
		ClientId:     "test",
	})
}

func TestOrderClient_MarketClose(t *testing.T) {
	call2(t, oc.MarketClose, OrderMarketCloseRequest{
		UserId:       102,
		Market:       "BTCUSDT",
		PositionId:   122,
		TakerFeeRate: "0.001",
		Amount:       "0",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		Option:       0,
		ClientId:     "okkk",
	})
}

func TestOrderClient_Cancel(t *testing.T) {
	call2(t, oc.Cancel, OrderCancelRequest{
		UserId:  userId1,
		Market:  "BTCUSDT",
		OrderId: 57,
	})
}

func TestOrderClient_CancelAll(t *testing.T) {
	call1(t, oc.CancelAll, OrderCancelAllRequest{
		UserId: userId1,
		Market: "BTCUSDT",
		Side:   Long,
	})
}

func TestOrderClient_Book(t *testing.T) {
	resp := call2(t, oc.Book, OrderBookRequest{
		Market: "BTCUSDT",
		Side:   Long,
		Offset: 0,
		Limit:  10,
	})

	foreach(t, resp.Records)
}

func TestOrderClient_Depth(t *testing.T) {
	v := call2(t, oc.Depth, OrderDepthRequest{
		Market:   "BTCUSDT",
		Limit:    10,
		Interval: "0",
	})

	foreach(t, v.Asks)
	foreach(t, v.Bids)
}

func TestOrderClient_Pending(t *testing.T) {
	resp := call2(t, oc.Pending, OrderPendingRequest{
		UserId: userId1,
		Market: nil,
		Side:   types.SideAll,
		Offset: 0,
		Limit:  10,
	})
	foreach(t, resp.Records)
}

func TestOrderClient_PendingDetail(t *testing.T) {
	call2(t, oc.PendingDetail, OrderPendingDetailRequest{
		Market:  "BTCUSDT",
		OrderId: 69,
	})
}

func TestOrderClient_Deals(t *testing.T) {
	call2(t, oc.Deals, OrderDealsRequest{
		UserId:  userId1,
		OrderId: 58,
		Offset:  0,
		Limit:   10,
	})
}

func TestOrderClient_Finished(t *testing.T) {
	resp := call2(t, oc.Finished, OrderFinishedRequest{
		UserId:    userId1,
		Market:    "BTCUSDT",
		Side:      0,
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
	foreach(t, resp.Records)
}

func TestOrderClient_FinishedDetail(t *testing.T) {
	call2(t, oc.FinishedDetail, OrderFinishedDetailRequest{
		UserId:  userId1,
		OrderId: 58,
	})
}

func TestOrderClient_GetUserPreference(t *testing.T) {
	call2(t, oc.GetUserPreference, OrderGetUserPreferenceRequest{
		UserId: 1000,
	})
}

func TestOrderClient_SetUserPreference(t *testing.T) {
	call1(t, oc.SetUserPreference, OrderSetUserPreferenceRequest{
		UserId:       1000,
		PositionMode: Hedge,
	})
}

func TestOrderClient_LimitAdd(t *testing.T) {
	call2(t, oc.LimitAdd, OrderLimitAddRequest{
		UserId:       userId2,
		Market:       "BTCUSDT",
		PositionId:   2,
		Amount:       "0.1",
		Price:        "30000",
		TakerFeeRate: "0.01",
		MakerFeeRate: "0.01",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		ClientId:     "test",
	})
}

func TestOrderClient_MarketAdd(t *testing.T) {
	call2(t, oc.MarketAdd, OrderMarketAddRequest{
		UserId:       userId1,
		Market:       "BTCUSDT",
		PositionId:   2,
		Amount:       "0.1",
		TakerFeeRate: "0.01",
		Source:       "test",
		FeeAsset:     nil,
		FeeDiscount:  nil,
		ClientId:     "test",
	})
}

func TestOrderClient_UserTrade(t *testing.T) {
	call2(t, oc.UserTrade, OrderUserTradeRequest{
		UserId: 201,
		Market: "BTCUSDT",
	})
}
