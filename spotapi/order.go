package spotapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2"
	"github.com/ashexchange/matchclient/v2/types"
)

type OrderClient interface {
	// OrderPutLimit 下限价委托单
	OrderPutLimit(ctx context.Context, req *OrderPutLimitRequest) (*OrderDetail, error)

	// OrderPutMarket 下市价委托单
	OrderPutMarket(ctx context.Context, req *OrderPutMarketRequest) (*OrderDetail, error)

	// OrderCancel  撤销委托单
	OrderCancel(ctx context.Context, req *OrderCancelRequest) (*OrderDetail, error)

	// OrderCancelBatch 批量撤销委托单
	OrderCancelBatch(ctx context.Context, req *OrderCancelBatchRequest) ([]*OrderCancelBatchResponse, error)

	// OrderCancelAll 撤销全部委托单
	OrderCancelAll(ctx context.Context, req *OrderCancelAllRequest) (*types.Result, error)

	// OrderCancelUserAll 撤销用户的全部委托单
	OrderCancelUserAll(ctx context.Context, req *OrderCancelUserAllRequest) (*types.Result, error)

	// OrderPutStopLimit 下计划限价委托单
	OrderPutStopLimit(ctx context.Context, req *OrderPutStopLimitRequest) (*types.Result, error)

	// OrderPutStopMarket 下市价计划委托单
	OrderPutStopMarket(ctx context.Context, req *OrderPutStopMarketRequest) (*types.Result, error)

	// OrderCancelStop 撤销计划委托单
	OrderCancelStop(ctx context.Context, req *OrderCancelStopRequest) (*StopOrderDetail, error)

	// OrderCancelStopAll 撤销全部计划委托单
	OrderCancelStopAll(ctx context.Context, req *OrderCancelStopAllRequest) (*types.Result, error)

	// OrderUserOrders 查找指定用户的，指定账号的订单
	OrderUserOrders(ctx context.Context, req *OrderUserOrdersRequest) (*OrderUserOrdersResponse, error)

	// OrderDeals 查询某个委托单的成交明细
	OrderDeals(ctx context.Context, req *OrderDealsRequest) (*OrderDeals, error)

	// OrderBook 查询某个市场的订单薄
	OrderBook(ctx context.Context, req *OrderBookRequest) (*OrderBookResponse, error)

	// OrderStopBook 查询某个市场的计划委托订单薄
	OrderStopBook(ctx context.Context, req *OrderStopBookRequest) (*StopOrderBookResponse, error)

	// OrderDepth 按照某个精度合并订单薄，既深度数据
	OrderDepth(ctx context.Context, req *OrderDepthRequest) (*OrderDepthResponse, error)

	// OrderPending 取得某个用户的当前委托单(及还处于市场深度中的委托单)
	OrderPending(ctx context.Context, req *OrderPendingRequest) (*OrderPendingResponse, error)

	// OrderPendingInTime 实时取得某个用户的当前委托单(及还处于市场深度中的委托单)
	OrderPendingInTime(ctx context.Context, req *OrderPendingInTimeRequest) (*OrderPendingResponse, error)
	OrderPendingInTimeNew(ctx context.Context, req *OrderPendingInTimeNewRequest) (*OrderPendingResponse, error)

	// OrderPendingStop 获取某个用户未被触发的计划委托单
	OrderPendingStop(ctx context.Context, req *OrderPendingStopRequest) (*OrderPendingStopResponse, error)

	// OrderPendingStopInTime 实时获取某个用户未被触发的计划委托单
	OrderPendingStopInTime(ctx context.Context, req *OrderPendingStopInTimeRequest) (*OrderPendingStopResponse, error)
	OrderPendingStopInTimeNew(ctx context.Context, req *OrderPendingStopInTimeNewRequest) (*OrderPendingStopResponse, error)

	// OrderPendingDetail 获取某个用户当前委托的详细信息
	OrderPendingDetail(ctx context.Context, req *OrderPendingDetailRequest) (*OrderDetail, error)

	// OrderFinished 获取某个用户已完结的订单
	OrderFinished(ctx context.Context, req *OrderFinishedRequest) (*OrderMatchResponse, error)

	// OrderFinishedStop 获取某个用户已完结的计划委托
	OrderFinishedStop(ctx context.Context, req *OrderFinishedStopRequest) (*OrderFinishedStopResponse, error)

	// OrderFinishedDetail 获取用户已经完成的订单的详细信息
	OrderFinishedDetail(ctx context.Context, req *OrderFinishedDetailRequest) (*MatchOrder, error)

	// OrderCloseOrderAll 撤销某个市场全部委托
	OrderCloseOrderAll(ctx context.Context, req *OrderCloseOrderAllRequest) (*types.Result, error)

	// OrderCloseStopOrderAll 撤销某个市场全部的计划委托
	OrderCloseStopOrderAll(ctx context.Context, req *OrderCloseStopOrderAllRequest) (*types.Result, error)
}

// PutOrderOption 下单的可选项
type PutOrderOption uint

const (
	// suggest use stock fee
	UseStockFee PutOrderOption = 1 << iota

	// suggest use money fee
	UseMoneyFee

	// unlimited min amount
	UnlimitedMinAmount

	// immediate or cancel order
	ImmediateOrCancelOrder

	// fill or kill order
	FillOrKillOrder
)

type OrderPutLimitRequest struct {
	UserId       types.UserID
	Account      int
	Market       string
	Side         types.Side
	Amount       string
	Price        string
	TakerFeeRate string
	MakerFeeRate string
	Source       string  // limit 30 bytes
	FeeAsset     *string // must be null or asset name

	// 0～1
	FeeDiscount string
	Option      PutOrderOption
	ClientId    string
}

type OrderDetail struct {
	Id          uint64     `json:"id"`
	Type        int        `json:"type"`
	Side        types.Side `json:"side"`
	User        int64      `json:"user"`
	Account     int        `json:"account"`
	Option      int        `json:"option"`
	Ctime       float64    `json:"ctime"`
	Mtime       float64    `json:"mtime"`
	Market      string     `json:"market"`
	Source      string     `json:"source"`
	ClientId    string     `json:"client_id"`
	Price       string     `json:"price"`
	Amount      string     `json:"amount"`
	TakerFee    string     `json:"taker_fee"`
	MakerFee    string     `json:"maker_fee"`
	Left        string     `json:"left"`
	DealStock   string     `json:"deal_stock"`
	DealMoney   string     `json:"deal_money"`
	MoneyFee    string     `json:"money_fee"`
	StockFee    string     `json:"stock_fee"`
	AssetFee    string     `json:"asset_fee"`
	FeeDiscount string     `json:"fee_discount"`
	FeeAsset    *string    `json:"fee_asset"`
	Status      int        `json:"status"`
	Finished    bool       `json:"finished"`
}

type OrderPutMarketRequest struct {
	UserId       types.UserID
	Account      int
	Market       string
	Side         types.Side
	Amount       string
	TakerFeeRate string
	Source       string  // limit 30 bytes
	FeeAsset     *string // must be null or asset name

	// 0～1
	FeeDiscount string

	Option   PutOrderOption
	ClientId string
}

type OrderCancelRequest struct {
	UserId  types.UserID
	Market  string
	OrderId uint64
}

type OrderCancelBatchRequest struct {
	UserId   types.UserID
	Market   string
	OrderIds []types.ID
}

type OrderCancelBatchResponse struct {
	Code    int64        `json:"code"`
	Message string       `json:"message"`
	Order   *OrderDetail `json:"order"`
}

type OrderCancelAllRequest struct {
	UserId  types.UserID
	Account int
	Market  string
	Side    types.Side // 可选值
}

type OrderCancelUserAllRequest struct {
	UserId  types.UserID
	Account int
}

type OrderPutStopLimitRequest struct {
	UserId       types.UserID
	Account      int
	Market       string
	Side         types.Side
	Amount       string
	StopPrice    string
	Price        string
	TakerFeeRate string
	MakerFeeRate string
	Source       string  // limit 30 bytes
	FeeAsset     *string // must be null or asset name

	// 0～1
	FeeDiscount string

	Option   PutOrderOption
	ClientId string
}

type OrderPutStopMarketRequest struct {
	UserId       types.UserID
	Account      int
	Market       string
	Side         types.Side
	Amount       string
	Price        string
	StopPrice    string
	TakerFeeRate string
	Source       string  // limit 30 bytes
	FeeAsset     *string // must be null or asset name

	// 0～1
	FeeDiscount string

	Option   PutOrderOption
	ClientId string
}

type OrderCancelStopRequest struct {
	UserId  types.UserID
	Market  string
	OrderId uint64
}

type OrderCancelStopAllRequest struct {
	UserId  types.UserID
	Account int
	Market  string
	Side    types.Side // 可选值
}

type OrderUserOrdersRequest struct {
	UserId  types.UserID
	Account int
	Market  string
	Side    types.Side // 可选值
	State   int
	Type    int
	Offset  int64
	Limit   int64
}

type OrderDealsRequest struct {
	UserId  types.UserID
	Account int
	OrderId uint64
	Offset  int64
	Limit   int64
}

type OrderDeals struct {
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	Records []*OrderDealsRecord
}

type OrderDealsRecord struct {
	DealUser    uint64  `json:"deal_user"`
	Account     int     `json:"account"`
	Fee         string  `json:"fee"`
	Deal        string  `json:"deal"`
	Price       string  `json:"price"`
	Amount      string  `json:"amount"`
	Role        int     `json:"role"` // 1：Maker, 2: Taker
	User        uint64  `json:"user"`
	Time        float64 `json:"time"`
	FeeAsset    *string `json:"fee_asset"`
	DealOrderId uint64  `json:"deal_order_id"`
	Id          uint64  `json:"id"`
}

type OrderBookRequest struct {
	Market      string
	Side        types.Side // 1: sell, 2: buy
	Offset      int64
	Limit       int64
	FilterUsers []types.UserID // 过滤用户
}

type OrderBookResponse struct {
	Offset int64          `json:"offset"`
	Limit  int64          `json:"limit"`
	Orders []*OrderDetail `json:"orders"`
}

type OrderStopBookRequest struct {
	Market      string
	State       int // 1: low, 2: high
	Offset      int64
	Limit       int64
	FilterUsers []types.UserID // 过滤用户
}

type StopOrderDetail struct {
	Id          uint64     `json:"id"`
	Type        int        `json:"type"`
	Side        types.Side `json:"side"`
	User        int64      `json:"user"`
	Account     int        `json:"account"`
	Option      int64      `json:"option"`
	State       int        `json:"state"`
	Ctime       float64    `json:"ctime"`
	Mtime       float64    `json:"mtime"`
	Market      string     `json:"market"`
	Source      string     `json:"source"`
	ClientId    string     `json:"client_id"`
	StopPrice   string     `json:"stop_price"`
	Price       string     `json:"price"`
	Amount      string     `json:"amount"`
	TakerFee    string     `json:"taker_fee"`
	MakerFee    string     `json:"maker_fee"`
	FeeDiscount string     `json:"fee_discount"`
	FeeAsset    *string    `json:"fee_asset"`
}

type OrderUserOrdersResponse struct{}

type StopOrderBookResponse struct {
	Limit  int64              `json:"limit"`
	Offset int64              `json:"offset"`
	Total  int64              `json:"total"`
	Orders []*StopOrderDetail `json:"orders"`
}

type OrderDepthRequest struct {
	Market   string
	Limit    int64
	Interval string
}

// 询价
type Inquiry [2]types.Number

type OrderDepthResponse struct {
	Asks     []*Inquiry            `json:"asks"`
	Bids     []*Inquiry            `json:"bids"`
	UpdateId uint64                `json:"update_id"`
	Last     string                `json:"last"`
	Time     types.TimeMillisecond `json:"time"`
}

type OrderPendingRequest struct {
	UserId  types.UserID
	Account int
	Market  *string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingResponse struct {
	Limit   int64          `json:"limit"`
	Offset  int64          `json:"offset"`
	Total   int64          `json:"total"`
	Records []*OrderDetail `json:"records"`
}

type OrderPendingInTimeRequest struct {
	UserId  types.UserID
	Account int
	Market  string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingInTimeNewRequest struct {
	UserId  types.UserID
	Account int
	Market  *string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingStopRequest struct {
	UserId  types.UserID
	Account int
	Market  *string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingStopResponse struct {
	Limit   int64              `json:"limit"`
	Offset  int64              `json:"offset"`
	Total   int64              `json:"total"`
	Records []*StopOrderDetail `json:"records"`
}

type OrderPendingStopInTimeRequest struct {
	UserId  types.UserID
	Account int
	Market  string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingStopInTimeNewRequest struct {
	UserId  types.UserID
	Account int
	Market  *string
	Side    types.Side
	Offset  int64
	Limit   int64
}

type OrderPendingDetailRequest struct {
	Market  string
	OrderId uint64
}

type OrderFinishedRequest struct {
	UserId    types.UserID
	Account   int
	Market    string
	Side      types.Side
	StartTime int64
	EndTime   int64
	Offset    int64
	Limit     int64
	Status    []int
	Option    int // 1 过滤无成交的撤单
}

type MatchOrder struct {
	Id          int64      `json:"id"`
	CTime       float64    `json:"ctime"`
	FTime       float64    `json:"ftime"`
	User        uint64     `json:"user"`
	Account     int        `json:"account"`
	Option      int        `json:"option"`
	Market      string     `json:"market"`
	Source      string     `json:"source"`
	Type        int        `json:"type"`
	Side        types.Side `json:"side"`
	Price       string     `json:"price"`
	Amount      string     `json:"amount"`
	TakerFee    string     `json:"taker_fee"`
	MakerFee    string     `json:"maker_fee"`
	DealStock   string     `json:"deal_stock"`
	DealMoney   string     `json:"deal_money"`
	MoneyFee    string     `json:"money_fee"`
	StockFee    string     `json:"stock_fee"`
	FeeAsset    *string    `json:"fee_asset"`
	FeeDiscount string     `json:"fee_discount"`
	AssetFee    string     `json:"asset_fee"`
	ClientId    string     `json:"client_id"`
	DealFlag    int        `json:"deal_flag"` // 成交标志 0-无成交 1-部分成交 2-完全成交
	Status      int        `json:"status"`    // 状态: 1-完成 2-错误 3-撤销
}
type OrderMatchResponse struct {
	Offset  int64         `json:"offset"`
	Limit   int64         `json:"limit"`
	Total   int64         `json:"total"`
	Records []*MatchOrder `json:"records"`
}

type OrderFinishedStopRequest struct {
	UserId    types.UserID
	Account   int
	Market    string
	Side      types.Side
	StartTime int64
	EndTime   int64
	Offset    int64
	Limit     int64
	Status    []int
}

type MatchStopOrder struct {
	Account     int        `json:"account"`
	FeeDiscount string     `json:"fee_discount"`
	StopPrice   string     `json:"stop_price"`
	CTime       float64    `json:"ctime"`
	MakerFee    string     `json:"maker_fee"`
	Price       string     `json:"price"`
	Side        types.Side `json:"side"`
	Source      string     `json:"source"`
	Amount      string     `json:"amount"`
	User        int64      `json:"user"`
	MTime       float64    `json:"mtime"`
	FeeAsset    *string    `json:"fee_asset"`
	Type        int        `json:"type"`
	Id          uint64     `json:"id"`
	RealOrderId uint64     `json:"real_order_id"` // 委托单触发后的id
	Market      string     `json:"market"`
	TakerFee    string     `json:"taker_fee"`
	ClientId    string     `json:"client_id"`
	Status      int        `json:"status"`
}

type OrderFinishedStopResponse struct {
	Offset  int64             `json:"offset"`
	Limit   int64             `json:"limit"`
	Total   int64             `json:"total"`
	Records []*MatchStopOrder `json:"records"`
}

type OrderFinishedDetailRequest struct {
	UserId  types.UserID
	OrderId uint64
}

type OrderCloseOrderAllRequest struct {
	Market string `json:"market"`
}

type OrderCloseStopOrderAllRequest struct {
	Market string `json:"market"`
}

type orderClient struct {
	cli matchclient.Invoker
}

func NewOrderClient(c matchclient.Invoker) OrderClient {
	return &orderClient{cli: c}
}

func (c *orderClient) OrderPutLimit(ctx context.Context, req *OrderPutLimitRequest) (reply *OrderDetail, err error) {
	err = c.cli.Invoke(ctx, "order.put_limit", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Amount,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)
	return
}

func (c *orderClient) OrderPutMarket(ctx context.Context, req *OrderPutMarketRequest) (reply *OrderDetail, err error) {
	err = c.cli.Invoke(ctx, "order.put_market", &reply,
		req.UserId,
		req.Account,
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

func (c *orderClient) OrderCancel(ctx context.Context, req *OrderCancelRequest) (reply *OrderDetail, err error) {
	err = c.cli.Invoke(ctx, "order.cancel", &reply, req.UserId, req.Market, req.OrderId)
	return
}

func (c *orderClient) OrderCancelBatch(ctx context.Context, req *OrderCancelBatchRequest) (reply []*OrderCancelBatchResponse, err error) {
	err = c.cli.Invoke(ctx, "order.cancel_batch", &reply, req.UserId, req.Market, req.OrderIds)
	return
}

func (c *orderClient) OrderCancelAll(ctx context.Context, req *OrderCancelAllRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.cancel_all", &reply, req.UserId, req.Account, req.Market, req.Side)
	return
}

func (c *orderClient) OrderCancelUserAll(ctx context.Context, req *OrderCancelUserAllRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.cancel_user_all", &reply, req.UserId, req.Account)
	return
}

func (c *orderClient) OrderPutStopLimit(ctx context.Context, req *OrderPutStopLimitRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.put_stop_limit", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Amount,
		req.StopPrice,
		req.Price,
		req.TakerFeeRate,
		req.MakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)
	return
}

func (c *orderClient) OrderPutStopMarket(ctx context.Context, req *OrderPutStopMarketRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.put_stop_market", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Amount,
		req.StopPrice,
		req.TakerFeeRate,
		req.Source,
		req.FeeAsset,
		req.FeeDiscount,
		req.Option,
		req.ClientId,
	)
	return
}

func (c *orderClient) OrderCancelStop(ctx context.Context, req *OrderCancelStopRequest) (reply *StopOrderDetail, err error) {
	err = c.cli.Invoke(ctx, "order.cancel_stop", &reply, req.UserId, req.Market, req.OrderId)
	return
}

func (c *orderClient) OrderCancelStopAll(ctx context.Context, req *OrderCancelStopAllRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.cancel_stop_all", &reply, req.UserId, req.Account, req.Market, req.Side)
	return
}

func (c *orderClient) OrderUserOrders(ctx context.Context, req *OrderUserOrdersRequest) (reply *OrderUserOrdersResponse, err error) {
	err = c.cli.Invoke(ctx, "order.user_orders", &reply, req.UserId, req.Account, req.Market, req.Side, req.State,
		req.Type, req.Offset, req.Limit)
	return
}

func (c *orderClient) OrderDeals(ctx context.Context, req *OrderDealsRequest) (reply *OrderDeals, err error) {
	err = c.cli.Invoke(ctx, "order.deals", &reply,
		req.UserId,
		req.Account,
		req.OrderId,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderBook(ctx context.Context, req *OrderBookRequest) (reply *OrderBookResponse, err error) {
	err = c.cli.Invoke(ctx, "order.book", &reply, req.Market, req.Side, req.Offset, req.Limit, req.FilterUsers)
	return
}

func (c *orderClient) OrderStopBook(ctx context.Context, req *OrderStopBookRequest) (reply *StopOrderBookResponse, err error) {
	err = c.cli.Invoke(ctx, "order.stop_book", &reply, req.Market, req.State, req.Offset, req.Limit, req.FilterUsers)
	return
}

func (c *orderClient) OrderDepth(ctx context.Context, req *OrderDepthRequest) (reply *OrderDepthResponse, err error) {
	err = c.cli.Invoke(ctx, "order.depth", &reply, req.Market, req.Limit, req.Interval)
	return
}

func (c *orderClient) OrderPending(ctx context.Context, req *OrderPendingRequest) (reply *OrderPendingResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingInTime(ctx context.Context, req *OrderPendingInTimeRequest) (reply *OrderPendingResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending_intime", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingInTimeNew(ctx context.Context, req *OrderPendingInTimeNewRequest) (reply *OrderPendingResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending_intime", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingStop(ctx context.Context, req *OrderPendingStopRequest) (reply *OrderPendingStopResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending_stop", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingStopInTime(ctx context.Context, req *OrderPendingStopInTimeRequest) (reply *OrderPendingStopResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending_stop_intime", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingStopInTimeNew(ctx context.Context, req *OrderPendingStopInTimeNewRequest) (reply *OrderPendingStopResponse, err error) {
	err = c.cli.Invoke(ctx, "order.pending_stop_intime", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *orderClient) OrderPendingDetail(ctx context.Context, req *OrderPendingDetailRequest) (reply *OrderDetail, err error) {
	err = c.cli.Invoke(ctx, "order.pending_detail", &reply, req.Market, req.OrderId)
	return
}

func (c *orderClient) OrderFinished(ctx context.Context, req *OrderFinishedRequest) (reply *OrderMatchResponse, err error) {
	err = c.cli.Invoke(ctx, "order.finished", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
		req.Status,
		req.Option,
	)
	return
}

func (c *orderClient) OrderFinishedStop(ctx context.Context, req *OrderFinishedStopRequest) (reply *OrderFinishedStopResponse, err error) {
	err = c.cli.Invoke(ctx, "order.finished_stop", &reply,
		req.UserId,
		req.Account,
		req.Market,
		req.Side,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
		req.Status,
	)
	return
}

func (c *orderClient) OrderFinishedDetail(ctx context.Context, req *OrderFinishedDetailRequest) (reply *MatchOrder, err error) {
	err = c.cli.Invoke(ctx, "order.finished_detail", &reply, req.UserId, req.OrderId)
	return
}

func (c *orderClient) OrderCloseOrderAll(ctx context.Context, req *OrderCloseOrderAllRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.close_order_all", &reply, req.Market)
	return
}

func (c *orderClient) OrderCloseStopOrderAll(ctx context.Context, req *OrderCloseStopOrderAllRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "order.close_stop_all", &reply, req.Market)
	return
}
