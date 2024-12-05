package futuresapi

import (
	"testing"
	"time"
)

var pc = NewPositionClient(cli)

func TestPositionClient_AdjustMargin(t *testing.T) {
	call2(t, pc.AdjustMargin, PositionAdjustMarginRequest{
		UserId:     userId1,
		Market:     "BTCUSDT",
		PositionId: 18,
		Type:       MarginSub,
		Amount:     "1",
	})
}

func TestPositionClient_Deals(t *testing.T) {
	call2(t, pc.Deals, PositionDealsRequest{
		UserId:     userId1,
		PositionId: 9,
		Offset:     0,
		Limit:      10,
	})
}

func TestPositionClient_Finished(t *testing.T) {
	call2(t, pc.Finished, PositionFinishedRequest{
		UserId:    userId1,
		Market:    "BTCUSDT",
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
}

func TestPositionClient_Funding(t *testing.T) {
	call2(t, pc.Funding, PositionFundingRequest{
		UserId:    userId1,
		Market:    "BTCUSDT",
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
		Offset:    0,
		Limit:     10,
	})
}

func TestPositionClient_LimitConfig(t *testing.T) {
	r, err := pc.LimitConfig(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range r {
		for _, i := range v {
			t.Logf("%+v", i)
		}
	}
}

func TestPositionClient_LiqingList(t *testing.T) {
	call2(t, pc.LiqingList, PositionLiqingListRequest{
		Market: "BTCUSDT",
		Side:   Short,
		Offset: 0,
		Limit:  10,
	})
}

func TestPositionClient_List(t *testing.T) {
	call2(t, pc.List, PositionListRequest{
		Market: "BTCUSDT",
		Side:   Long,
		Offset: 0,
		Limit:  10,
	})
}

func TestPositionClient_Margins(t *testing.T) {
	resp := call2(t, pc.Margins, PositionMarginsRequest{
		UserId:     userId1,
		PositionId: 9,
		Offset:     0,
		Limit:      10,
	})
	foreach(t, resp.Records)
}

func TestPositionClient_Pending(t *testing.T) {
	resp := call2(t, pc.Pending, PositionPendingRequest{
		UserId: userId1,
		Market: nil,
	})
	foreach(t, resp)
}

func TestPositionClient_AdjustLeverage(t *testing.T) {
	call1(t, pc.AdjustLeverage, PositionAdjustLeverageRequest{
		UserId:       userId2,
		Market:       "BTCUSDT",
		PositionId:   7,
		PositionType: Isolated,
		Leverage:     "10",
	})
}

func TestPositionClient_View(t *testing.T) {
	call2(t, pc.View, PositionViewRequest{
		UserId:     1001,
		Market:     "BTCUSDT",
		PositionId: 7,
	})
}
