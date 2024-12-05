package spotapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/matchclient/v2"
	"github.com/ashexchange/matchclient/v2/types"
)

type MarketClient interface {
	// MarketLast 获取某个市场的最新成交价
	MarketLast(ctx context.Context, req *MarketLastRequest) (string, error)

	// MarketDetail 获取市场列表
	MarketList(ctx context.Context) ([]*MarketDetail, error)

	// MarketDetail 获取市场详细信息
	MarketDetail(ctx context.Context, req *MarketDetailRequest) (*MarketDetail, error)

	// MarketDeals 获取当前市场的最新成交
	MarketDeals(ctx context.Context, req *MarketDealsRequest) ([]*MarketDeal, error)

	// MarketSummary 获取市场统计信息
	MarketSummary(ctx context.Context, req *MarketSummaryRequest) (*MarketSummary, error)

	// MarketDealsExt 获取市场最新成交详情
	MarketDealsExt(ctx context.Context, req *MarketDealsExtRequest) ([]*MarketDealsExt, error)

	// MarketUserDeals 获取某个用户的交易记录
	MarketUserDeals(ctx context.Context, req *MarketUserDealsRequest) (*MarketUserDealResponse, error)

	// MarketKline 获取市场k线
	MarketKline(ctx context.Context, req *MarketKlineRequest) (MarketKlineResult, error)

	// MarketStatus 获取市场状态
	MarketStatus(ctx context.Context, req *MarketStatusRequest) (*MarketStatus, error)

	// MarketSelfDeal 自成交，不经过撮合(用来刷量跟画k线)
	MarketSelfDeal(ctx context.Context, req *MarketSelfDealRequest) (*types.Result, error)
}

type MarketLastRequest struct {
	Market string
}

type MarketDetail struct {
	Name      string `json:"name"`
	Stock     string `json:"stock"`
	Money     string `json:"money"`
	Account   int    `json:"account"`
	FeePrec   int    `json:"fee_prec"`
	StockPrec int    `json:"stock_prec"`
	MoneyPrec int    `json:"money_prec"`
	MinAmount string `json:"min_amount"`
}

type MarketDetailRequest struct {
	Market string
}

type MarketDealsRequest struct {
	Market string
	Limit  int64
	LastId int64
}

type MarketSummaryRequest struct {
	Market string
}

type MarketDeal struct {
	Id     int64   `json:"id"`
	Time   float64 `json:"time"`
	Type   string  `json:"type"`
	Amount string  `json:"amount"`
	Price  string  `json:"price"`
}

type MarketSummary struct {
	OrderUsers     int64  `json:"order_users"`
	OrderAskUsers  int64  `json:"order_ask_users"`
	OrderBidUsers  int64  `json:"order_bid_users"`
	StopUsers      int64  `json:"stop_users"`
	StopAskUsers   int64  `json:"stop_ask_users"`
	StopBidUsers   int64  `json:"stop_bid_users"`
	Orders         int64  `json:"orders"`
	Stops          int64  `json:"stops"`
	OrderAsks      int64  `json:"order_asks"`
	OrderAskAmount string `json:"order_ask_amount"`
	OrderAskLeft   string `json:"order_ask_left"`
	OrderBids      string `json:"order_bids"`
	OrderBidAmount string `json:"order_bid_amount"`
	OrderBidLeft   string `json:"order_bid_left"`
	StopAsks       string `json:"stop_asks"`
	StopAskAmount  string `json:"stop_ask_amount"`
	StopBids       string `json:"stop_bids"`
	StopBidAmount  string `json:"stop_bid_amount"`
}

type MarketDealsExtRequest struct {
	Market string
	Limit  int64
	LastId int64
}

type MarketDealsExt struct {
	Id        int64   `json:"id"`
	Time      float64 `json:"time"`
	Type      string  `json:"type"`
	Price     string  `json:"price"`
	Amount    string  `json:"amount"`
	AskUserId uint64  `json:"ask_user_id"`
	BidUserId uint64  `json:"bid_user_id"`
}

type MarketUserDealsRequest struct {
	UserId    types.UserID
	AccountId int // 0 is compatible with the original, -1 for query all
	Market    string
	Side      types.Side // 0 for no limit, 1 for sell, 2 for buy
	StartTime int64
	EndTime   int64
	Offset    int64
	Limit     int64
}

type MarketUserDealResponse struct {
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	Total   int64 `json:"total"`
	Records []*MarketUserDeal
}

type MarketUserDeal struct {
	Id           uint64     `json:"id"`
	OrderId      uint64     `json:"order_id"`
	Time         float64    `json:"time"`
	User         uint64     `json:"user"`
	Market       string     `json:"market"`
	Account      int        `json:"account"`
	DealUser     uint64     `json:"deal_user"`
	Side         types.Side `json:"side"` // 1: sell, 2: buy
	Role         int        `json:"role"` // 1: Maker, 2: Taker
	Price        string     `json:"price"`
	Amount       string     `json:"amount"` // 数量
	Deal         string     `json:"deal"`   // 总价值 price * amount
	Fee          string     `json:"fee"`
	FeeAsset     string     `json:"fee_asset"`
	DealAccount  int        `json:"deal_account"`
	DealOrderId  uint64     `json:"deal_order_id"`
	DealFee      string     `json:"deal_fee"`
	DealFeeAsset string     `json:"deal_fee_asset"`
}

type MarketKlineRequest struct {
	Market   string // 市场
	Start    uint64 // 开始时间
	End      uint64 // 结束时间
	Interval uint   // kline类型
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

type MarketStatusRequest struct {
	Market string // market name and get index status by market name with suffix _INDEX like: BTCUSDT_INDEX
	Period uint   // e.g. 86400 for last 24 hours
}

type MarketStatus struct {
	Period    uint   `json:"period"`
	Last      string `json:"last"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	SellTotal string `json:"sell_total"`
	BuyTotal  string `json:"buy_total"`
	Deal      string `json:"deal"`
}

type MarketSelfDealRequest struct {
	Market string
	Amount string
	Price  string
	Side   types.Side
}

type marketClient struct {
	cli matchclient.Invoker
}

func NewMarketClient(c matchclient.Invoker) MarketClient {
	return &marketClient{cli: c}
}

func (c *marketClient) MarketLast(ctx context.Context, req *MarketLastRequest) (reply string, err error) {
	err = c.cli.Invoke(ctx, "market.last", &reply, req.Market)
	return
}

func (c *marketClient) MarketList(ctx context.Context) (reply []*MarketDetail, err error) {
	err = c.cli.Invoke(ctx, "market.list", &reply)
	return
}

func (c *marketClient) MarketDetail(ctx context.Context, req *MarketDetailRequest) (reply *MarketDetail, err error) {
	err = c.cli.Invoke(ctx, "market.detail", &reply, req.Market)
	return
}

func (c *marketClient) MarketDeals(ctx context.Context, req *MarketDealsRequest) (reply []*MarketDeal, err error) {
	err = c.cli.Invoke(ctx, "market.deals", &reply, req.Market, req.Limit, req.LastId)
	return
}

func (c *marketClient) MarketSummary(ctx context.Context, req *MarketSummaryRequest) (reply *MarketSummary, err error) {
	err = c.cli.Invoke(ctx, "market.summary", &reply, req.Market)
	return
}

func (c *marketClient) MarketDealsExt(ctx context.Context, req *MarketDealsExtRequest) (reply []*MarketDealsExt, err error) {
	err = c.cli.Invoke(ctx, "market.deals_ext", &reply, req.Market, req.Limit, req.LastId)
	return
}

func (c *marketClient) MarketUserDeals(ctx context.Context, req *MarketUserDealsRequest) (reply *MarketUserDealResponse, err error) {
	err = c.cli.Invoke(ctx, "market.user_deals", &reply,
		req.UserId, req.AccountId, req.Market, req.Side, req.StartTime, req.EndTime, req.Offset, req.Limit,
	)
	return
}

func (c *marketClient) MarketKline(ctx context.Context, req *MarketKlineRequest) (reply MarketKlineResult, err error) {
	err = c.cli.Invoke(ctx, "market.kline", &reply, req.Market, req.Start, req.End, req.Interval)
	return
}

func (c *marketClient) MarketStatus(ctx context.Context, req *MarketStatusRequest) (reply *MarketStatus, err error) {
	err = c.cli.Invoke(ctx, "market.status", &reply, req.Market, req.Period)
	return
}

func (c *marketClient) MarketSelfDeal(ctx context.Context, req *MarketSelfDealRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "market.self_deal", &reply, req.Market, req.Amount, req.Price, req.Side)
	return
}
