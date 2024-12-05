package futuresapi

import (
	"testing"
	"time"

	"github.com/ashexchange/matchclient/v2/types"
)

var mc = NewMarketClient(cli)

func TestMarketClient_AdjustLeverage(t *testing.T) {
	call2(t, mc.AdjustLeverage, MarketAdjustLeverageRequest{
		UserId:       userId1,
		Market:       "BTCUSDT",
		PositionType: Isolated,
		Leverage:     "10",
	})
	call2(t, mc.AdjustLeverage, MarketAdjustLeverageRequest{
		UserId:       userId2,
		Market:       "BTCUSDT",
		PositionType: Isolated,
		Leverage:     "10",
	})
}

func TestMarketClient_Deals(t *testing.T) {
	r := call2(t, mc.Deals, MarketDealsRequest{
		Market: "BTCUSDT",
		Limit:  10,
		LastId: 0,
	})

	foreach(t, r)
}

func TestMarketClient_DealsExt(t *testing.T) {
	resp := call2(t, mc.DealsExt, MarketDealsExtRequest{
		Market: "BTCUSDT",
		Limit:  10,
		LastId: 0,
	})

	foreach(t, resp)
}

func TestMarketClient_FundingHistory(t *testing.T) {
	resp := call2(t, mc.FundingHistory, MarketFundingHistoryRequest{
		Market:    "BTCUSDT",
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
	foreach(t, resp.Records)
}

func TestMarketClient_GetPreference(t *testing.T) {
	call2(t, mc.GetPreference, MarketGetPreferenceRequest{
		UserId: userId1,
		Market: "BTCUSDT",
	})
}

func TestMarketClient_Insurances(t *testing.T) {
	call0(t, mc.Insurances)
}

func TestMarketClient_Kline(t *testing.T) {
	call2(t, mc.Kline, MarketKlineRequest{
		Market:    "BTCUSDT",
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Interval:  60 * 60,
	})
}

func TestMarketClient_Last(t *testing.T) {
	call2(t, mc.Last, MarketLastRequest{Market: "BTCUSDT"})
}

func TestMarketClient_LimitConfig(t *testing.T) {
	v, _ := mc.LimitConfig(ctx)
	for _, l := range v {
		for _, i := range l {
			t.Logf("%s", i)
		}
	}
}

func TestMarketClient_List(t *testing.T) {
	resp := call0(t, mc.List)
	foreach(t, resp)
}

func TestMarketClient_PremiumHistory(t *testing.T) {
	call2(t, mc.PremiumHistory, MarketPremiumHistoryRequest{
		Market:    "BTCUSDT",
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
}

func TestMarketClient_SelfDeal(t *testing.T) {
	call1(t, mc.SelfDeal, MarketSelfDealRequest{
		Market: "BTCUSDT",
		Side:   Long,
		Amount: "10",
		Price:  "20000",
	})
}

func TestMarketClient_Status(t *testing.T) {
	call2(t, mc.Status, MarketStatusRequest{
		Market: "BTCUSDT",
		Period: 60,
	})
}

func TestMarketClient_Summary(t *testing.T) {
	call2(t, mc.Summary, MarketSummaryRequest{
		Market: "BTCUSDT",
	})
}

func TestMarketClient_UserDeals(t *testing.T) {
	call2(t, mc.UserDeals, MarketUserDealsRequest{
		UserId:    userId1,
		Market:    "BTCUSDT",
		Side:      Long,
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
}

func TestMarketClient_MakerQuery(t *testing.T) {
	call2(t, mc.MakerQuery, MarketMakerQueryRequest{
		Market: "BTCUSDT",
	})
}

func TestMarketClient_MakerUpdate(t *testing.T) {
	call1(t, mc.MakerUpdate, MarketMakerUpdateRequest{
		Market: "BTCUSDT",
		Op:     MakerUpdateOpAdd,
		Users: []types.UserID{
			200,
			201,
			202,
		},
	})
	call1(t, mc.MakerUpdate, MarketMakerUpdateRequest{
		Market: "BTCUSDT",
		Op:     MakerUpdateOpDel,
		Users: []types.UserID{
			201,
		},
	})
}

func TestMarketClient_UserTrade(t *testing.T) {
	call2(t, mc.UserTrade, MarketUserTradeRequest{
		Market: "BTCUSDT",
	})
}

func TestMarketClient_AllUserTrade(t *testing.T) {
	res, err := mc.AllUserTrade(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}
