package futuresapi

import (
	"testing"
	"time"

	"github.com/ashexchange/matchclient/v2/types"
)

var tc = NewTradeClient(cli)

func TestTradeClient_AmountRank(t *testing.T) {
	call2(t, tc.AmountRank, TradeAmountRankRequest{
		Markets:   []string{"BTCUSDT"},
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
	})
}

func TestTradeClient_DealSummary(t *testing.T) {
	call2(t, tc.DealSummary, TradeDealSummaryRequest{
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
	})
}

func TestTradeClient_NetRank(t *testing.T) {
	call2(t, tc.NetRank, TradeNetRankRequest{
		Markets:   []string{"BTCUSDT"},
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
	})
}

func TestTradeClient_UsersVolume(t *testing.T) {
	call2(t, tc.UsersVolume, TradeUsersVolumeRequest{
		Markets:   []string{"BTCUSDT"},
		Users:     []types.UserID{1000},
		StartTime: time.Now().Unix() - 24*60*60,
		EndTime:   time.Now().Unix(),
	})
}
