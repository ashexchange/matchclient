package futuresapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/kit"
	"github.com/ashexchange/matchclient/v2/types"
)

type AssetClient struct {
	invoker Invoker
}

func NewAssetClient(invoker Invoker) AssetClient {
	return AssetClient{invoker}
}

type AssetInfo struct {
	Name      string `json:"name"`
	Precision int    `json:"prec"`
}

type AssetInfoResult []*AssetInfo

func (c AssetClient) List(ctx context.Context) (result AssetInfoResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.list", &result)

	return
}

type AssetUpdateRequest struct {
	UserId     types.UserID
	Asset      string
	Business   string
	BusinessId types.ID
	Change     string
	Detail     map[string]any
}

var noDetails = struct{}{}

func (c AssetClient) Update(ctx context.Context, req AssetUpdateRequest) error {
	return c.invoker.Invoke(ctx, "asset.update", nil,
		req.UserId,
		req.Asset,
		req.Business,
		req.BusinessId,
		req.Change,
		kit.Ternary[any](req.Detail == nil, noDetails, req.Detail),
	)
}

type AssetQueryRequest struct {
	UserId types.UserID
	Assets []string
}

type AssetBalanceDetail struct {
	Available    types.Number `json:"available"`
	Frozen       types.Number `json:"frozen"`
	Margin       types.Number `json:"margin"`
	BalanceTotal types.Number `json:"balance_total"`
	ProfitUnreal types.Number `json:"profit_unreal"`
	Transfer     types.Number `json:"transfer"`
}

func (bd *AssetBalanceDetail) String() string {
	return fmt.Sprintf("{Availabe: %s, Frozen: %s, Margin: %s, BalanceTotal: %s, ProfitUnreal: %s, Transfer: %s}",
		bd.Available,
		bd.Frozen,
		bd.Margin,
		bd.BalanceTotal,
		bd.ProfitUnreal,
		bd.Transfer,
	)
}

func (bd *AssetBalanceDetail) clone() *AssetBalanceDetail {
	b := *bd
	return &b
}

type AssetQueryResult map[string]*AssetBalanceDetail

var zeroBalance = AssetBalanceDetail{
	Available:    "0",
	Frozen:       "0",
	Margin:       "0",
	BalanceTotal: "0",
	ProfitUnreal: "0",
	Transfer:     "0",
}

func (r AssetQueryResult) Get(name string) *AssetBalanceDetail {
	if bd := r[name]; bd != nil {
		return bd
	}

	return zeroBalance.clone()
}

func (c AssetClient) Query(ctx context.Context, req AssetQueryRequest) (result AssetQueryResult, err error) {
	params := make([]any, 1, len(req.Assets)+1)
	params[0] = req.UserId
	for _, asset := range req.Assets {
		params = append(params, asset)
	}

	err = c.invoker.Invoke(ctx, "asset.query", &result, params...)

	return
}

func (c AssetClient) QueryAsset(ctx context.Context, userId types.UserID, asset string) (*AssetBalanceDetail, error) {
	var result AssetQueryResult
	err := c.invoker.Invoke(ctx, "asset.query", &result, userId, asset)
	return result.Get(asset), err
}

type AssetQueryUsersRequest struct {
	Asset string
	Users []types.UserID
}

type AssetQueryUsersResult map[types.UserID]*AssetBalanceDetail

func (r AssetQueryUsersResult) Get(userId types.UserID) *AssetBalanceDetail {
	if bd := r[userId]; bd != nil {
		return bd
	}

	return zeroBalance.clone()
}

func (c AssetClient) QueryUsers(ctx context.Context, req AssetQueryUsersRequest) (result AssetQueryUsersResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.query_users", &result,
		req.Asset,
		req.Users,
	)

	return
}

type AssetSummaryRequest struct {
	Asset string
}

type AssetSummaryResult struct {
	Transfer *AssetSummaryTransfer `json:"transfer"`
	Asset    *AssetSummaryDetail   `json:"asset"`
	Time     types.Time            `json:"time"`
}

type AssetSummaryDetail struct {
	UserTotal   types.Number `json:"user_total"`
	Insurance   types.Number `json:"insurance"`
	FeeTotal    types.Number `json:"fee_total"`
	ProfitTotal types.Number `json:"profit_total"`
}

type AssetSummaryTransfer struct {
	In  types.Number `json:"in"`
	Out types.Number `json:"out"`
}

func (c AssetClient) Summary(ctx context.Context, req AssetSummaryRequest) (result *AssetSummaryResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.summary", &result, req.Asset)

	return
}

type AssetFeeAssetSummaryRequest struct {
	Asset string
}

type AssetFeeAssetSummaryResult struct {
	Transfer *AssetSummaryTransfer       `json:"transfer"`
	Asset    *AssetFeeAssetSummaryDetail `json:"asset"`
	Time     types.Time                  `json:"time"`
}

type AssetFeeAssetSummaryDetail struct {
	UserTotal types.Number `json:"user_total"`
	FeeTotal  types.Number `json:"fee_total"`
}

func (c AssetClient) FeeAssetSummary(ctx context.Context, req AssetFeeAssetSummaryRequest) (result *AssetFeeAssetSummaryResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.fee_asset_summary", &result, req.Asset)

	return
}

type AssetHistoryRequest struct {
	UserId    types.UserID
	Asset     string
	Business  string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type AssetHistoryResult struct {
	Records []*AssetHistory `json:"records"`
	Offset  int32           `json:"offset"`
	Limit   int32           `json:"limit"`
}

type AssetHistory struct {
	UserId   types.UserID   `json:"user_id"`
	Detail   map[string]any `json:"detail"`
	Time     types.Time     `json:"time"`
	Asset    string         `json:"asset"`
	Business string         `json:"business"`
	Change   types.Number   `json:"change"`
	Balance  types.Number   `json:"balance"`
}

func (c AssetClient) History(ctx context.Context, req AssetHistoryRequest) (result *AssetHistoryResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.history", &result,
		req.UserId,
		req.Asset,
		req.Business,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type AssetHistoryAllRequest struct {
	UserId    types.UserID
	Asset     string
	Business  string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type AssetHistoryAllResult struct {
	Records []*AssetHistory `json:"records"`
	Offset  int32           `json:"offset"`
	Limit   int32           `json:"limit"`
}

func (c AssetClient) HistoryAll(ctx context.Context, req AssetHistoryAllRequest) (result *AssetHistoryAllResult, err error) {
	err = c.invoker.Invoke(ctx, "asset.history_all", &result,
		req.UserId,
		req.Asset,
		req.Business,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}
