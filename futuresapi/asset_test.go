package futuresapi

import (
	"math/rand"
	"testing"

	"github.com/ashexchange/matchclient/v2/types"
)

var ac = NewAssetClient(cli)

func TestAssetClient_List(t *testing.T) {
	resp := call0(t, ac.List)
	foreach(t, resp)
}

func TestAssetClient_Update(t *testing.T) {
	call1(t, ac.Update, AssetUpdateRequest{
		UserId:     userId1,
		Asset:      "USDT",
		Business:   "test",
		BusinessId: uint64(rand.Int()),
		Change:     "100000.0",
		Detail:     map[string]any{"ok": "ok"},
	})
}

func TestAssetClient_Query(t *testing.T) {
	t.Log(ac.QueryAsset(ctx, 100, "USDT"))
}

func TestAssetClient_QueryUsers(t *testing.T) {
	call2(t, ac.QueryUsers, AssetQueryUsersRequest{
		Asset: "USDT",
		Users: []types.UserID{1001, 1000},
	})
}

func TestAssetClient_Summary(t *testing.T) {
	call2(t, ac.Summary, AssetSummaryRequest{Asset: "USDT"})
}

func TestAssetClient_FeeAssetSummary(t *testing.T) {
	call2(t, ac.FeeAssetSummary, AssetFeeAssetSummaryRequest{Asset: "USDT"})
}

func TestAssetClient_History(t *testing.T) {
	call2(t, ac.History, AssetHistoryRequest{
		UserId:    userId1,
		Asset:     "USDT",
		Business:  "test",
		StartTime: 0,
		EndTime:   0,
		Offset:    0,
		Limit:     10,
	})
}

func TestAssetClient_HistoryAll(t *testing.T) {
	call2(t, ac.HistoryAll, AssetHistoryAllRequest{
		UserId:    userId1,
		Asset:     "USDT",
		Business:  "test",
		StartTime: 0,
		EndTime:   0,
		Offset:    0,
		Limit:     10,
	})
}
