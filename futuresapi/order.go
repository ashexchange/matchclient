package futuresapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/matchclient/v2/types"
)

type OrderClient struct {
	invoker Invoker
}

func NewOrderClient(invoker Invoker) OrderClient {
	return OrderClient{invoker}
}

func optionEffectType(option OrderOption, effectType OrderEffectType) uint64 {
	if effectType == 0 {
		effectType = EffectDefault
	}

	return uint64(option)<<32 | uint64(effectType)
}

type OrderDetail struct {
	Market         string          `json:"market"`
	OrderId        types.ID        `json:"order_id"`
	UserId         types.UserID    `json:"user_id"`
	Target         Target          `json:"target"`
	Side           types.Side      `json:"side"`
	Direction      Direction       `json:"direction"`
	Type           types.OrderType `json:"type"`
	Left           types.Number    `json:"left"`
	Amount         types.Number    `json:"amount"`
	Price          types.Number    `json:"price"`
	PositionId     types.ID        `json:"position_id"`
	PositionType   PositionType    `json:"position_type"`
	Leverage       types.Number    `json:"leverage"`
	EffectType     OrderEffectType `json:"effect_type"`
	Source         string          `json:"source"`
	DealProfit     types.Number    `json:"deal_profit"`
	DealAssetFee   types.Number    `json:"deal_asset_fee"`
	DealFee        types.Number    `json:"deal_fee"`
	MakerFee       types.Number    `json:"maker_fee"`
	TakerFee       types.Number    `json:"taker_fee"`
	FeeDiscount    types.Number    `json:"fee_discount"`
	FeeAsset       string          `json:"fee_asset"`
	DealStock      types.Number    `json:"deal_stock"`
	ClientId       string          `json:"client_id"`
	UpdateTime     types.Time      `json:"update_time"`
	CreateTime     types.Time      `json:"create_time"`
	LastDealId     types.ID        `json:"last_deal_id"`
	LastDealType   DealType        `json:"last_deal_type"`
	LastDealRole   types.TradeRole `json:"last_deal_role"`
	LastDealAmount types.Number    `json:"last_deal_amount"`
	LastDealPrice  types.Number    `json:"last_deal_price"`
	LastDealTime   types.Time      `json:"last_deal_time"`
}

type OrderPutLimitRequest struct {
	UserId       types.UserID
	Market       string
	Side         types.Side
	Amount       string
	Price        string
	TakerFeeRate string
	MakerFeeRate string
	// Source up to 30 bytes
	Source string
	// FeeAsset null or CET
	FeeAsset    *string
	FeeDiscount *string
	Option      OrderOption
	EffectType  OrderEffectType
	// ClientId
	// user self-define order id
	//
	// Optional: true
	ClientId string

	// optional
	PositionType PositionType
	Leverage     *string
}

func (c OrderClient) PutLimit(ctx context.Context, req OrderPutLimitRequest) (result *OrderDetail, err error) {
	if req.PositionType.IsValid() || req.Leverage != nil {
		return c.putLimitWithPositionType(ctx, req)
	}

	return c.putLimit(ctx, req)
}

func (c OrderClient) putLimit(ctx context.Context, req OrderPutLimitRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.put_limit", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.Amount,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		optionEffectType(req.Option, req.EffectType),
		req.ClientId,
	)

	return
}

func (c OrderClient) putLimitWithPositionType(ctx context.Context, req OrderPutLimitRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.put_limit", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.Amount,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		optionEffectType(req.Option, req.EffectType),
		req.ClientId,
		req.PositionType,
		req.Leverage,
	)

	return
}

type OrderPutMarketRequest struct {
	UserId       types.UserID
	Market       string
	Side         types.Side
	Amount       string
	TakerFeeRate string
	// Source up to 30 bytes
	Source string
	// FeeAsset null or CET
	FeeAsset    *string
	FeeDiscount *string
	Option      OrderOption
	// ClientId
	// user self-define order id
	//
	// Optional: true
	ClientId string

	// optional
	PositionType PositionType
	Leverage     *string
}

func (c OrderClient) PutMarket(ctx context.Context, req OrderPutMarketRequest) (result *OrderDetail, err error) {
	if req.PositionType.IsValid() || req.Leverage != nil {
		return c.putMarketWithPositionType(ctx, req)
	}

	return c.putMarket(ctx, req)
}

func (c OrderClient) putMarket(ctx context.Context, req OrderPutMarketRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.put_market", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.Amount,
		req.TakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)

	return
}

func (c OrderClient) putMarketWithPositionType(ctx context.Context, req OrderPutMarketRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.put_market", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.Amount,
		req.TakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
		req.PositionType,
		req.Leverage,
	)

	return
}

type OrderLimitCloseRequest struct {
	UserId       types.UserID
	Market       string
	PositionId   types.ID
	Amount       string
	Price        string
	TakerFeeRate string
	MakerFeeRate string
	// Source up to 30 bytes
	Source string
	// FeeAsset null or CET
	FeeAsset    *string
	FeeDiscount *string
	Option      OrderOption
	EffectType  OrderEffectType
	// ClientId
	// user self-define order id
	//
	// Optional: true
	ClientId string
}

func (c OrderClient) LimitClose(ctx context.Context, req OrderLimitCloseRequest) (result *OrderDetail, err error) {
	if req.Amount == "" {
		req.Amount = "0"
	}

	err = c.invoker.Invoke(ctx, "order.limit_close", &result,
		req.UserId,
		req.Market,
		req.PositionId,
		req.Amount,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		optionEffectType(req.Option, req.EffectType),
		req.ClientId,
	)

	return
}

type OrderMarketCloseRequest struct {
	UserId     types.UserID
	Market     string
	PositionId types.ID
	// Amount close all if the amount is zero
	Amount       string
	TakerFeeRate string
	// Source up to 30 bytes
	Source string
	// FeeAsset null or CET
	FeeAsset    *string
	FeeDiscount *string
	Option      OrderOption
	// ClientId
	// user self-define order id
	//
	// Optional: true
	ClientId string
}

func (c OrderClient) MarketClose(ctx context.Context, req OrderMarketCloseRequest) (result *OrderDetail, err error) {
	if req.Amount == "" {
		req.Amount = "0"
	}

	err = c.invoker.Invoke(ctx, "order.market_close", &result,
		req.UserId,
		req.Market,
		req.PositionId,
		req.Amount,
		req.TakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)

	return
}

type OrderCancelRequest struct {
	UserId  types.UserID
	Market  string
	OrderId types.ID
}

func (c OrderClient) Cancel(ctx context.Context, req OrderCancelRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.cancel", &result,
		req.UserId,
		req.Market,
		req.OrderId,
	)

	return
}

type OrderCancelAllRequest struct {
	UserId types.UserID
	Market string
	Side   types.Side
}

func (c OrderClient) CancelAll(ctx context.Context, req OrderCancelAllRequest) error {
	return c.invoker.Invoke(ctx, "order.cancel_all", nil,
		req.UserId,
		req.Market,
		req.Side,
	)
}

type OrderCancelBatchRequest struct {
	UserId types.UserID
	Market string
	// OrderIds lengths max is 101
	OrderIds []types.ID
}

type OrderCancelOrderResult struct {
	Code    int32        `json:"code"`
	Message string       `json:"message"`
	Order   *OrderDetail `json:"order"`
}

type OrderCancelBatchResult []*OrderCancelOrderResult

func (c OrderClient) CancelBatch(ctx context.Context, req OrderCancelBatchRequest) (result OrderCancelBatchResult, err error) {
	err = c.invoker.Invoke(ctx, "order.cancel_batch", &result,
		req.UserId,
		req.Market,
		req.OrderIds,
	)

	return
}

type OrderBookRequest struct {
	Market string
	Side   types.Side
	Offset int32
	Limit  int32
}

type OrderBookResult struct {
	Records []*OrderDetail `json:"records"`
	Offset  int32          `json:"offset"`
	Limit   int32          `json:"limit"`
	Total   int32          `json:"total"`
}

func (c OrderClient) Book(ctx context.Context, req OrderBookRequest) (result *OrderBookResult, err error) {
	err = c.invoker.Invoke(ctx, "order.book", &result,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)

	return
}

type OrderDepthRequest struct {
	Market string
	Limit  int32
	// Interval
	//
	// "0" for no interval
	Interval string
}

type OrderDepthResult struct {
	IndexPrice types.Number          `json:"index_price"`
	SignPrice  types.Number          `json:"sign_price"`
	Time       types.TimeMillisecond `json:"time"`
	Last       types.Number          `json:"last"`
	Asks       []DepthItem           `json:"asks"`
	Bids       []DepthItem           `json:"bids"`
}

type DepthItem [2]types.Number

func (item DepthItem) String() string {
	return fmt.Sprintf("DepthItem{Price:%s, Amount:%s}", item.Price(), item.Amount())
}

func (item DepthItem) SetPrice(s string)    { item[0] = types.Number(s) }
func (item DepthItem) SetAmount(s string)   { item[1] = types.Number(s) }
func (item DepthItem) Price() types.Number  { return item[0] }
func (item DepthItem) Amount() types.Number { return item[1] }

func (c OrderClient) Depth(ctx context.Context, req OrderDepthRequest) (result *OrderDepthResult, err error) {
	err = c.invoker.Invoke(ctx, "order.depth", &result,
		req.Market,
		req.Limit,
		req.Interval,
	)

	return
}

type OrderPendingRequest struct {
	UserId types.UserID
	// Market
	//
	// Null for all market
	Market *string
	// Side 0 for all side
	Side   types.Side
	Offset int32
	Limit  int32
}

type OrderPendingResult struct {
	Records []*OrderDetail `json:"records"`
	Offset  int32          `json:"offset"`
	Limit   int32          `json:"limit"`
	Total   int32          `json:"total"`
}

func (c OrderClient) Pending(ctx context.Context, req OrderPendingRequest) (result *OrderPendingResult, err error) {
	err = c.invoker.Invoke(ctx, "order.pending", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)

	return
}

type OrderPendingDetailRequest struct {
	Market  string
	OrderId types.ID
}

func (c OrderClient) PendingDetail(ctx context.Context, req OrderPendingDetailRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.pending_detail", &result,
		req.Market,
		req.OrderId,
	)

	return
}

type OrderDealsRequest struct {
	UserId  types.UserID
	OrderId types.ID
	Offset  int32
	Limit   int32
}

type OrderDealsResult struct {
	Records []*OrderDealDetail `json:"records"`
	Offset  int32              `json:"offset"`
	Limit   int32              `json:"limit"`
}

type OrderDealDetail struct {
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

func (c OrderClient) Deals(ctx context.Context, req OrderDealsRequest) (result *OrderDealsResult, err error) {
	err = c.invoker.Invoke(ctx, "order.deals", &result,
		req.UserId,
		req.OrderId,
		req.Offset,
		req.Limit,
	)

	return
}

type OrderFinishedRequest struct {
	UserId types.UserID
	// Market
	//
	// "" for all markets
	Market string
	// Side
	//
	// 0: all
	// 1: sell
	// 2: buy
	Side      types.Side
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type OrderFinishedResult struct {
	Records []*OrderFinishedDetail `json:"records"`
	Offset  int32                  `json:"offset"`
	Limit   int32                  `json:"limit"`
	Total   int32                  `json:"total"`
}

func (c OrderClient) Finished(ctx context.Context, req OrderFinishedRequest) (result *OrderFinishedResult, err error) {
	err = c.invoker.Invoke(ctx, "order.finished", &result,
		req.UserId,
		req.Market,
		req.Side,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type OrderFinishedDetailRequest struct {
	UserId  types.UserID
	OrderId types.ID
}

type OrderFinishedDetail struct {
	MakerFee     types.Number    `json:"maker_fee"`
	Source       string          `json:"source"`
	CreateTime   types.Time      `json:"create_time"`
	OrderId      types.ID        `json:"order_id"`
	PositionId   types.ID        `json:"position_id"`
	UpdateTime   types.Time      `json:"update_time"`
	Leverage     types.Number    `json:"leverage"`
	UserId       types.UserID    `json:"user_id"`
	Target       Target          `json:"target"`
	PositionType PositionType    `json:"position_type"`
	Market       string          `json:"market"`
	DealFee      types.Number    `json:"deal_fee"`
	TakerFee     types.Number    `json:"taker_fee"`
	Type         types.OrderType `json:"type"`
	EffectType   OrderEffectType `json:"effect_type"`
	Side         types.Side      `json:"side"`
	FeeAsset     string          `json:"fee_asset"`
	Amount       types.Number    `json:"amount"`
	DealProfit   types.Number    `json:"deal_profit"`
	Price        types.Number    `json:"price"`
	Left         types.Number    `json:"left"`
	DealStock    types.Number    `json:"deal_stock"`
	FeeDiscount  types.Number    `json:"fee_discount"`
	DealAssetFee types.Number    `json:"deal_asset_fee"`
	ClientId     string          `json:"client_id"`
}

func (c OrderClient) FinishedDetail(ctx context.Context, req OrderFinishedDetailRequest) (result *OrderFinishedDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.finished_detail", &result,
		req.UserId,
		req.OrderId,
	)

	return
}

type OrderGetUserPreferenceRequest struct {
	UserId types.UserID
}

type OrderGetUserPreferenceResult struct {
	PositionMode PositionMode `json:"position_mode"`
}

func (c OrderClient) GetUserPreference(ctx context.Context, req OrderGetUserPreferenceRequest) (result *OrderGetUserPreferenceResult, err error) {
	err = c.invoker.Invoke(ctx, "order.get_user_preference", &result, req.UserId)

	return
}

type OrderSetUserPreferenceRequest struct {
	UserId       types.UserID
	PositionMode PositionMode
}

func (c OrderClient) SetUserPreference(ctx context.Context, req OrderSetUserPreferenceRequest) (err error) {
	err = c.invoker.Invoke(ctx, "order.set_user_preference", nil,
		req.UserId,
		req.PositionMode,
	)

	return
}

type OrderLimitAddRequest struct {
	UserId       types.UserID
	Market       string
	PositionId   types.ID
	Amount       string
	Price        string
	TakerFeeRate string
	MakerFeeRate string
	Source       string
	FeeAsset     *string
	FeeDiscount  *string
	Option       OrderOption
	EffectType   OrderEffectType
	ClientId     string
}

func (c OrderClient) LimitAdd(ctx context.Context, req OrderLimitAddRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.limit_add", &result,
		req.UserId,
		req.Market,
		req.PositionId,
		req.Amount,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		optionEffectType(req.Option, req.EffectType),
		req.ClientId,
	)

	return
}

type OrderMarketAddRequest struct {
	UserId       types.UserID
	Market       string
	PositionId   types.ID
	Amount       string
	TakerFeeRate string
	Source       string
	FeeAsset     *string
	FeeDiscount  *string
	Option       OrderOption
	ClientId     string
}

func (c OrderClient) MarketAdd(ctx context.Context, req OrderMarketAddRequest) (result *OrderDetail, err error) {
	err = c.invoker.Invoke(ctx, "order.market_add", &result,
		req.UserId,
		req.Market,
		req.PositionId,
		req.Amount,
		req.TakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)

	return
}

type OrderUserTradeRequest struct {
	UserId types.UserID
	Market string
}

type OrderUserTradeResult struct {
	AskCount     int `json:"ask_count"`
	BidCount     int `json:"bid_count"`
	PositionMode int `json:"position_mode"`
	ShortCount   int `json:"short_count"`
	LongCount    int `json:"long_count"`
}

func (c OrderClient) UserTrade(ctx context.Context, req OrderUserTradeRequest) (result *OrderUserTradeResult, err error) {
	err = c.invoker.Invoke(ctx, "order.user_trade", &result,
		req.UserId,
		req.Market,
	)

	return
}
