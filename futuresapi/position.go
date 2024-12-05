package futuresapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/matchclient/v2/types"
)

type PositionClient struct {
	invoker Invoker
}

func NewPositionClient(invoker Invoker) PositionClient {
	return PositionClient{invoker}
}

type PositionPendingRequest struct {
	UserId types.UserID
	// Market null for all markets
	Market *string
}

type PositionDetail struct {
	UpdateTime              types.Time   `json:"update_time"`
	ProfitUnreal            types.Number `json:"profit_unreal"`
	OpenMargin              types.Number `json:"open_margin"`
	PositionId              types.ID     `json:"position_id"`
	CreateTime              types.Time   `json:"create_time"`
	LiqAmount               types.Number `json:"liq_amount"`
	Market                  string       `json:"market"`
	MaintenanceMarginAmount types.Number `json:"mainten_margin_amount"`
	ProfitClearing          types.Number `json:"profit_clearing"`
	FinishType              FinishType   `json:"finish_type"`
	UserId                  types.UserID `json:"user_id"`
	Type                    PositionType `json:"type"`
	Side                    types.Side   `json:"side"`
	MaintenanceMargin       types.Number `json:"mainten_margin"`
	Mode                    PositionMode `json:"mode"`
	Sys                     int          `json:"sys"`
	Amount                  types.Number `json:"amount"`
	AmountMax               types.Number `json:"amount_max"`
	LiqOrderPrice           types.Number `json:"liq_order_price"`
	CloseLeft               types.Number `json:"close_left"`
	LiqBrkDirection         BrkDirection `json:"liq_brk_direction"`
	OpenValMax              types.Number `json:"open_val_max"`
	LiqTime                 types.Time   `json:"liq_time"`
	Leverage                types.Number `json:"leverage"`
	OpenPrice               types.Number `json:"open_price"`
	AdlSortVal              types.Number `json:"adl_sort_val"`
	OpenVal                 types.Number `json:"open_val"`
	FeeAsset                string       `json:"fee_asset"`
	OpenMarginImply         types.Number `json:"open_margin_imply"`
	MarginAmount            types.Number `json:"margin_amount"`
	ProfitReal              types.Number `json:"profit_real"`
	LiqOrderTime            types.Time   `json:"liq_order_time"`
	LiqProfit               types.Number `json:"liq_profit"`
	DealAssetFee            types.Number `json:"deal_asset_fee"`
	LiqPrice                types.Number `json:"liq_price"`
	BkrPrice                types.Number `json:"bkr_price"`
	LiqPriceImply           types.Number `json:"liq_price_imply"`
	BkrPriceImply           types.Number `json:"bkr_price_imply"`
	AdlSort                 int          `json:"adl_sort"`
	Total                   int32        `json:"total"`
}

func (p PositionDetail) IsIsolated() bool {
	return p.Type.IsIsolated()
}

func (p PositionDetail) IsCross() bool {
	return p.Type.IsCross()
}

type PositionPendingResult []*PositionDetail

func (c PositionClient) Pending(ctx context.Context, req PositionPendingRequest) (result PositionPendingResult, err error) {
	err = c.invoker.Invoke(ctx, "position.pending", &result,
		req.UserId,
		req.Market,
	)

	return
}

type PositionListRequest struct {
	Market string
	Side   types.Side
	Offset int32
	Limit  int32
}

type PositionListResult struct {
	Records []*PositionDetail `json:"records"`
	Limit   int32             `json:"limit"`
	Offset  int32             `json:"offset"`
	Total   int32             `json:"total"`
}

func (c PositionClient) List(ctx context.Context, req PositionListRequest) (result *PositionListResult, err error) {
	err = c.invoker.Invoke(ctx, "position.list", &result,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionLiqingListRequest struct {
	Market string
	Side   types.Side
	Offset int32
	Limit  int32
}

type PositionLiqingListResult struct {
	Records []*PositionDetail `json:"records"`
	Limit   int32             `json:"limit"`
	Offset  int32             `json:"offset"`
	Total   int32             `json:"total"`
}

func (c PositionClient) LiqingList(ctx context.Context, req PositionLiqingListRequest) (result *PositionLiqingListResult, err error) {
	err = c.invoker.Invoke(ctx, "position.liqing_list", &result,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionAdjustMarginRequest struct {
	UserId     types.UserID
	Market     string
	PositionId types.ID
	Type       PositionAdjustMarginType
	Amount     string
}

func (c PositionClient) AdjustMargin(ctx context.Context, req PositionAdjustMarginRequest) (result *PositionDetail, err error) {
	err = c.invoker.Invoke(ctx, "position.adjust_margin", &result,
		req.UserId,
		req.Market,
		req.PositionId,
		req.Type,
		req.Amount,
	)

	return
}

type PositionAdjustLeverageRequest struct {
	UserId       types.UserID
	Market       string
	PositionId   types.ID
	PositionType PositionType
	Leverage     string
}

func (c PositionClient) AdjustLeverage(ctx context.Context, req PositionAdjustLeverageRequest) error {
	return c.invoker.Invoke(ctx, "position.adjust_leverage", nil,
		req.UserId,
		req.Market,
		req.PositionId,
		req.PositionType,
		req.Leverage,
	)
}

type PositionFinishedRequest struct {
	UserId    types.UserID
	Market    string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type PositionFinishedResult struct {
	Records []*PositionFinishedDetail `json:"records"`
	Offset  int32                     `json:"offset"`
	Limit   int32                     `json:"limit"`
	Total   int32                     `json:"total"`
}

type PositionFinishedDetail struct {
	Side              types.Side   `json:"side"`
	Sys               int          `json:"sys"`
	UpdateTime        types.Time   `json:"update_time"`
	PositionId        types.ID     `json:"position_id"`
	CreateTime        types.Time   `json:"create_time"`
	DealAssetFee      types.Number `json:"deal_asset_fee"`
	OpenPrice         types.Number `json:"open_price"`
	LiqProfit         types.Number `json:"liq_profit"`
	Leverage          types.Number `json:"leverage"`
	UserId            types.UserID `json:"user_id"`
	Market            string       `json:"market"`
	FinishType        FinishType   `json:"finish_type"`
	LiqAmount         types.Number `json:"liq_amount"`
	Type              PositionType `json:"type"`
	AmountMax         types.Number `json:"amount_max"`
	OpenValMax        types.Number `json:"open_val_max"`
	LiqPrice          types.Number `json:"liq_price"`
	BkrPrice          types.Number `json:"bkr_price"`
	MaintenanceMargin types.Number `json:"mainten_margin"`
	ProfitReal        types.Number `json:"profit_real"`
	FeeAsset          string       `json:"fee_asset"`
}

func (c PositionClient) Finished(ctx context.Context, req PositionFinishedRequest) (result *PositionFinishedResult, err error) {
	err = c.invoker.Invoke(ctx, "position.finished", &result,
		req.UserId,
		req.Market,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionDealsRequest struct {
	UserId     types.UserID
	PositionId types.ID
	Offset     int32
	Limit      int32
}

type PositionDealsResult struct {
	Records []*PositionDealDetail `json:"records"`
	Offset  int32                 `json:"offset"`
	Limit   int32                 `json:"limit"`
}

type PositionDealDetail struct {
	DealMargin     types.Number    `json:"deal_margin"`
	DealOrderId    types.ID        `json:"deal_order_id"`
	Id             types.ID        `json:"id"`
	PositionType   PositionType    `json:"position_type"`
	Market         string          `json:"market"`
	DealUserId     types.UserID    `json:"deal_user_id"`
	PositionAmount types.Number    `json:"position_amount"`
	Time           types.Time      `json:"time"`
	Side           types.Side      `json:"side"`
	DealType       DealType        `json:"deal_type"`
	FeePrice       types.Number    `json:"fee_price"`
	OpenPrice      types.Number    `json:"open_price"`
	FeeRate        types.Number    `json:"fee_rate"`
	Leverage       types.Number    `json:"leverage"`
	UserId         types.UserID    `json:"user_id"`
	OrderId        types.ID        `json:"order_id"`
	PositionId     types.ID        `json:"position_id"`
	MarginAmount   types.Number    `json:"margin_amount"`
	Role           types.TradeRole `json:"role"`
	FeeAsset       string          `json:"fee_asset"`
	Amount         types.Number    `json:"amount"`
	DealProfit     types.Number    `json:"deal_profit"`
	Price          types.Number    `json:"price"`
	DealFee        types.Number    `json:"deal_fee"`
	DealInsurance  types.Number    `json:"deal_insurance"`
	DealStock      types.Number    `json:"deal_stock"`
	FeeDiscount    types.Number    `json:"fee_discount"`
	FeeRealRate    types.Number    `json:"fee_real_rate"`
}

func (c PositionClient) Deals(ctx context.Context, req PositionDealsRequest) (result *PositionDealsResult, err error) {
	err = c.invoker.Invoke(ctx, "position.deals", &result,
		req.UserId,
		req.PositionId,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionMarginsRequest struct {
	UserId     types.UserID
	PositionId types.ID
	Offset     int32
	Limit      int32
}

type PositionMarginsResult struct {
	Records []*PositionMarginDetail `json:"records"`
	Offset  int32                   `json:"offset"`
	Limit   int32                   `json:"limit"`
}

type PositionMarginDetail struct {
	Time         types.Time   `json:"time"`
	Market       string       `json:"market"`
	UserId       types.UserID `json:"user_id"`
	PositionId   types.ID     `json:"position_id"`
	MarginChange types.Number `json:"margin_change"`
	// Type
	// 1: 增
	// 2: 减
	Type         int          `json:"type"`
	MarginAmount types.Number `json:"margin_amount"`
	LiqPrice     types.Number `json:"liq_price"`
	BkrPrice     types.Number `json:"bkr_price"`
}

func (c PositionClient) Margins(ctx context.Context, req PositionMarginsRequest) (result *PositionMarginsResult, err error) {
	err = c.invoker.Invoke(ctx, "position.margins", &result,
		req.UserId,
		req.PositionId,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionFundingRequest struct {
	UserId types.UserID
	// Market
	// "" for all markets
	Market    string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type PositionFundingResult struct {
	Records []*PositionFundingDetail `json:"records"`
	Offset  int32                    `json:"offset"`
	Limit   int32                    `json:"limit"`
}

type PositionFundingDetail struct {
	Price       types.Number `json:"price"`
	Amount      types.Number `json:"amount"`
	UserId      types.UserID `json:"user_id"`
	Funding     types.Number `json:"funding"`
	FundingRate types.Number `json:"funding_rate"`
	Time        types.Time   `json:"time"`
	Asset       string       `json:"asset"`
	Market      string       `json:"market"`
	// Type
	// 1: 支付
	// 2: 收取
	Type       int        `json:"type"`
	Side       types.Side `json:"side"`
	PositionId types.ID   `json:"position_id"`
}

func (c PositionClient) Funding(ctx context.Context, req PositionFundingRequest) (result *PositionFundingResult, err error) {
	err = c.invoker.Invoke(ctx, "position.funding", &result,
		req.UserId,
		req.Market,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type PositionLimitConfigItem [4]types.Number

func (m PositionLimitConfigItem) String() string {
	return fmt.Sprintf("LimitConfig{MinAmount:%s, MaxAmount:%s, Leverage:%s, MaintenanceMarginRate:%s}",
		m.MinAmount(),
		m.MaxAmount(),
		m.Leverage(),
		m.MaintenanceMarginRate(),
	)
}

func (m PositionLimitConfigItem) MinAmount() types.Number             { return m[0] }
func (m PositionLimitConfigItem) MaxAmount() types.Number             { return m[1] }
func (m PositionLimitConfigItem) Leverage() types.Number              { return m[2] }
func (m PositionLimitConfigItem) MaintenanceMarginRate() types.Number { return m[3] }

type PositionLimitConfigResult map[string][]PositionLimitConfigItem

func (c PositionClient) LimitConfig(ctx context.Context) (result PositionLimitConfigResult, err error) {
	err = c.invoker.Invoke(ctx, "position.limit_config", &result)

	return
}

type PositionViewRequest struct {
	UserId     types.UserID
	Market     string
	PositionId types.ID
}

func (c PositionClient) View(ctx context.Context, req PositionViewRequest) (result *PositionDetail, err error) {
	err = c.invoker.Invoke(ctx, "position.view", &result,
		req.UserId,
		req.Market,
		req.PositionId,
	)

	return
}
