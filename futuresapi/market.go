package futuresapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/matchclient/v2/types"
)

type MarketClient struct {
	invoker Invoker
}

func NewMarketClient(invoker Invoker) MarketClient {
	return MarketClient{invoker}
}

type MarketAdjustLeverageRequest struct {
	UserId       types.UserID
	Market       string
	PositionType PositionType
	Leverage     string
}

type MarketAdjustLeverageResult struct {
	Leverage     types.Number `json:"leverage"`
	PositionType PositionType `json:"position_type"`
}

// AdjustLeverage set the user default leverage and position type.
func (c MarketClient) AdjustLeverage(ctx context.Context, req MarketAdjustLeverageRequest) (result *MarketAdjustLeverageResult, err error) {
	err = c.invoker.Invoke(ctx, "market.adjust_leverage", &result,
		req.UserId,
		req.Market,
		req.PositionType,
		req.Leverage,
	)

	return
}

type MarketLimitConfigItem [3]types.Number

func (m MarketLimitConfigItem) String() string {
	return fmt.Sprintf("LimitConfig{Amount:%s, Leverage:%s, MaintenanceMarginRate:%s}",
		m.Amount(),
		m.Leverage(),
		m.MaintenanceMarginRate(),
	)
}

func (m MarketLimitConfigItem) Amount() types.Number                { return m[0] }
func (m MarketLimitConfigItem) Leverage() types.Number              { return m[1] }
func (m MarketLimitConfigItem) MaintenanceMarginRate() types.Number { return m[2] }

type LimitConfigResult map[string][]MarketLimitConfigItem

func (c MarketClient) LimitConfig(ctx context.Context) (result LimitConfigResult, err error) {
	err = c.invoker.Invoke(ctx, "market.limit_config", &result)

	return
}

type MarketGetPreferenceRequest struct {
	UserId types.UserID
	Market string
}

type MarketGetPreferenceResult struct {
	Leverage     types.Number `json:"leverage"`
	PositionType PositionType `json:"position_type"`
}

func (c MarketClient) GetPreference(ctx context.Context, req MarketGetPreferenceRequest) (result *MarketGetPreferenceResult, err error) {
	err = c.invoker.Invoke(ctx, "market.get_preference", &result,
		req.UserId,
		req.Market,
	)

	return
}

type MarketLastRequest struct {
	Market string
}

func (c MarketClient) Last(ctx context.Context, req MarketLastRequest) (result string, err error) {
	err = c.invoker.Invoke(ctx, "market.last", &result,
		req.Market,
	)

	return
}

type MarketInfo struct {
	Name            string         `json:"name"`
	StockPrecision  int            `json:"stock_prec"`
	Available       bool           `json:"available"`
	Type            MarketType     `json:"type"`
	MoneyPrecision  int            `json:"money_prec"`
	Leverages       []types.Number `json:"leverages"`
	Money           string         `json:"money"`
	Stock           string         `json:"stock"`
	FeePrecision    int            `json:"fee_prec"`
	TickSize        types.Number   `json:"tick_size"`
	Multiplier      types.Number   `json:"multiplier"`
	AmountPrecision int            `json:"amount_prec"`
	AmountMin       types.Number   `json:"amount_min"`
}

type MarketInfoResult []*MarketInfo

func (c MarketClient) List(ctx context.Context) (result MarketInfoResult, err error) {
	err = c.invoker.Invoke(ctx, "market.list", &result)

	return
}

type MarketSummaryRequest struct {
	Market string
}

type MarketSummaryResult struct {
	Time                types.Time   `json:"time"`
	PositionShortUsers  int          `json:"position_short_users"`
	PositionLongAmount  types.Number `json:"position_long_amount"`
	PositionLongUsers   int          `json:"position_long_users"`
	PositionShortAmount types.Number `json:"position_short_amount"`
}

func (c MarketClient) Summary(ctx context.Context, req MarketSummaryRequest) (result *MarketSummaryResult, err error) {
	err = c.invoker.Invoke(ctx, "market.summary", &result,
		req.Market,
	)

	return
}

type MarketInsurance struct {
	Time   types.Time   `json:"time"`
	Asset  string       `json:"asset"`
	Amount types.Number `json:"amount"`
}

func (c MarketClient) Insurances(ctx context.Context) (result []*MarketInsurance, err error) {
	err = c.invoker.Invoke(ctx, "market.insurances", &result)

	return
}

type MarketDealsRequest struct {
	Market string
	// Limit greater than 0, no more than 10000
	Limit  int32
	LastId int
}

type MarketDealsResult []*MarketDeal

type MarketDeal struct {
	Id     types.ID     `json:"id"`
	Price  types.Number `json:"price"`
	Amount types.Number `json:"amount"`
	Type   string       `json:"type"`
	Time   types.Time   `json:"time"`
}

func (c MarketClient) Deals(ctx context.Context, req MarketDealsRequest) (result MarketDealsResult, err error) {
	err = c.invoker.Invoke(ctx, "market.deals", &result,
		req.Market,
		req.Limit,
		req.LastId,
	)

	return
}

type MarketSelfDealRequest struct {
	Market string
	Side   types.Side
	Amount string
	Price  string
}

func (c MarketClient) SelfDeal(ctx context.Context, req MarketSelfDealRequest) error {
	return c.invoker.Invoke(ctx, "market.self_deal", nil,
		req.Market,
		req.Side,
		req.Amount,
		req.Price,
	)
}

type MarketDealsExtRequest struct {
	Market string
	// Limit no more than 10000
	Limit  int32
	LastId int
}

type MarketDealsExtResult []*MarketDealExt

type MarketDealExt struct {
	Type      string       `json:"type"`
	Id        types.ID     `json:"id"`
	Side      types.Side   `json:"side"`
	Amount    types.Number `json:"amount"`
	Time      types.Time   `json:"time"`
	AskUserId types.UserID `json:"ask_user_id"`
	BidUserId types.UserID `json:"bid_user_id"`
	Price     types.Number `json:"price"`
}

func (c MarketClient) DealsExt(ctx context.Context, req MarketDealsExtRequest) (result MarketDealsExtResult, err error) {
	err = c.invoker.Invoke(ctx, "market.deals_ext", &result,
		req.Market,
		req.Limit,
		req.LastId,
	)

	return
}

type MarketStatusRequest struct {
	Market string
	// Period cycle period, e.g. 86400 for last 24 hours
	Period int
}

type MarketStatusResult struct {
	Last               types.Number `json:"last"`
	Period             int          `json:"period"`
	Insurance          types.Number `json:"insurance"`
	FundingTime        int          `json:"funding_time"`
	FundingRateNext    types.Number `json:"funding_rate_next"`
	PositionAmount     types.Number `json:"position_amount"`
	IndexPrice         types.Number `json:"index_price"`
	FundingRateLast    types.Number `json:"funding_rate_last"`
	FundingRatePredict types.Number `json:"funding_rate_predict"`
	Close              types.Number `json:"close"`
	SignPrice          types.Number `json:"sign_price"`
	Open               types.Number `json:"open"`
	High               types.Number `json:"high"`
	Deal               types.Number `json:"deal"`
	Low                types.Number `json:"low"`
	Volume             types.Number `json:"volume"`
	SellTotal          types.Number `json:"sell_total"`
	BuyTotal           types.Number `json:"buy_total"`
}

func (c MarketClient) Status(ctx context.Context, req MarketStatusRequest) (result *MarketStatusResult, err error) {
	err = c.invoker.Invoke(ctx, "market.status", &result,
		req.Market,
		req.Period,
	)

	return
}

type MarketKlineRequest struct {
	Market    string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Interval  int
}

type MarketKlineResult []MarketCandle

type MarketCandle [8]any

func (c MarketCandle) String() string {
	return fmt.Sprintf("MarketCandle{Time:%f, Open:%s, Close:%s, High:%s, Low:%s, Volume:%s, Amount:%s, Market:%s}",
		c.Time(),
		c.Open(),
		c.Close(),
		c.High(),
		c.Low(),
		c.Volume(),
		c.Amount(),
		c.Market(),
	)
}

func (c MarketCandle) Time() types.Time     { return types.Time(c[0].(float64)) }
func (c MarketCandle) Open() types.Number   { return types.Number(c[1].(string)) }
func (c MarketCandle) Close() types.Number  { return types.Number(c[2].(string)) }
func (c MarketCandle) High() types.Number   { return types.Number(c[3].(string)) }
func (c MarketCandle) Low() types.Number    { return types.Number(c[4].(string)) }
func (c MarketCandle) Volume() types.Number { return types.Number(c[5].(string)) }
func (c MarketCandle) Amount() types.Number { return types.Number(c[6].(string)) }
func (c MarketCandle) Market() string       { return c[7].(string) }

func (c MarketClient) Kline(ctx context.Context, req MarketKlineRequest) (result MarketKlineResult, err error) {
	err = c.invoker.Invoke(ctx, "market.kline", &result,
		req.Market,
		req.StartTime,
		req.EndTime,
		req.Interval,
	)

	return
}

type MarketUserDealsRequest struct {
	UserId types.UserID
	// Market "" for all market
	Market string
	// Side 0: no limit, 1: sell, 2: buy
	Side      types.Side
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type MarketUserDealsResult struct {
	Records []*MarketUserDeal `json:"records"`
	Offset  int32             `json:"offset"`
	Limit   int32             `json:"limit"`
}

type MarketUserDeal struct {
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

func (c MarketClient) UserDeals(ctx context.Context, req MarketUserDealsRequest) (result *MarketUserDealsResult, err error) {
	err = c.invoker.Invoke(ctx, "market.user_deals", &result,
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

type MarketPremiumHistoryRequest struct {
	Market    string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type MarketPremiumHistoryResult struct {
	Records []*MarketPremiumHistory `json:"records"`
	Offset  int32                   `json:"offset"`
	Limit   int32                   `json:"limit"`
}

type MarketPremiumHistory struct {
	Time    types.Time   `json:"time"`
	Premium types.Number `json:"premium"`
	Market  string       `json:"market"`
}

func (c MarketClient) PremiumHistory(ctx context.Context, req MarketPremiumHistoryRequest) (result *MarketPremiumHistoryResult, err error) {
	err = c.invoker.Invoke(ctx, "market.premium_history", &result,
		req.Market,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type MarketFundingHistoryRequest struct {
	Market    string
	StartTime types.Timestamp
	EndTime   types.Timestamp
	Offset    int32
	Limit     int32
}

type FundingHistoryResult struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Records []*FundingHistory `json:"records"`
}

type FundingHistory struct {
	FundingRate     string       `json:"funding_rate"`
	Time            types.Time   `json:"time"`
	Asset           string       `json:"asset"`
	Market          string       `json:"market"`
	FundingRateReal types.Number `json:"funding_rate_real"`
}

func (c MarketClient) FundingHistory(ctx context.Context, req MarketFundingHistoryRequest) (result *FundingHistoryResult, err error) {
	err = c.invoker.Invoke(ctx, "market.funding_history", &result,
		req.Market,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)

	return
}

type MakerUpdateOp int

const (
	MakerUpdateOpAdd MakerUpdateOp = iota + 1
	MakerUpdateOpDel
)

type MarketMakerUpdateRequest struct {
	Market string
	Op     MakerUpdateOp
	Users  []types.UserID
}

func (c MarketClient) MakerUpdate(ctx context.Context, req MarketMakerUpdateRequest) error {
	return c.invoker.Invoke(ctx, "market.maker_update", nil,
		req.Market,
		req.Op,
		req.Users,
	)
}

type MarketMakerQueryRequest struct {
	Market string
}

type MarketMakerQueryResult []types.UserID

func (c MarketClient) MakerQuery(ctx context.Context, req MarketMakerQueryRequest) (result MarketMakerQueryResult, err error) {
	err = c.invoker.Invoke(ctx, "market.maker_query", &result,
		req.Market,
	)

	return
}

type MarketUserTradeRequest struct {
	Market string
}

type MarketUserTradeResult struct {
	BidCount     int                              `json:"bid_count"`
	AskCount     int                              `json:"ask_count"`
	UserCount    int                              `json:"user_count"`  // 有仓位用户数
	LongCount    int                              `json:"long_count"`  // 总多仓
	ShortCount   int                              `json:"short_count"` // 总空仓
	UserPosition []MarketUserTradePositionSummary `json:"user_position"`
}

type MarketUserTradePositionSummary struct {
	UserId       types.UserID `json:"user_id"`
	BidCount     int          `json:"bid_count"`
	AskCount     int          `json:"ask_count"`
	LongCount    int          `json:"long_count"`
	PositionMode int          `json:"position_mode"`
	ShortCount   int          `json:"short_count"`
}

func (c MarketClient) UserTrade(ctx context.Context, req MarketUserTradeRequest) (result *MarketUserTradeResult, err error) {
	err = c.invoker.Invoke(ctx, "market.user_trade", &result,
		req.Market,
	)

	return
}

type MarketAllUserTradeResult struct {
	Position   []MarketAllUserTradePositionSummary `json:"position"`
	LongCount  int                                 `json:"long_count"`
	ShortCount int                                 `json:"short_count"`
}

type MarketAllUserTradePositionSummary struct {
	Market     string `json:"market"`
	LongCount  int    `json:"long_count"`
	ShortCount int    `json:"short_count"`
}

func (c MarketClient) AllUserTrade(ctx context.Context) (result *MarketAllUserTradeResult, err error) {
	err = c.invoker.Invoke(ctx, "market.all_user_trade", &result)

	return
}
