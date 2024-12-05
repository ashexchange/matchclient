package spotapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2"
	"github.com/ashexchange/matchclient/v2/types"
)

type TradeClient interface {
	// TradeNetRank 净成交排行榜
	TradeNetRank(ctx context.Context, req *TradeNetRankRequest) (*TradeNetRankResponse, error)

	// TradeAmountRank 成交量排行榜
	TradeAmountRank(ctx context.Context, req *TradeAmountRankRequest) (*TradeAmountRankResponse, error)

	// TradeDealSummary 全站成交人数统计
	TradeDealSummary(ctx context.Context, req *TradeDealSummaryRequest) (*TradeDealSummaryResponse, error)

	// 交易额信息
	TradeUsersVolume(ctx context.Context, req *TradeUsersVolumeRequest) (TradeUsersVolumeResponse, error)
}

type TradeNetRankRequest struct {
	Market    string `json:"market"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type NetRank struct {
	Net    string
	Total  string
	UserId types.UserID
}
type TradeNetRankResponse struct {
	Sell           []*NetRank `json:"sell"`
	Buy            []*NetRank `json:"buy"`
	TotalAmount    string     `json:"total_amount"`
	TotalNet       string     `json:"total_net"`
	TotalSellUsers int        `json:"total_sell_users"`
	TotalBuyUsers  int        `json:"total_buy_users"`
}

type TradeAmountRankRequest struct {
	Market    []string `json:"market"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

type TradeAmount struct {
	Amount      string       `json:"amount"`
	UserId      types.UserID `json:"user_id"`
	TotalAmount string       `json:"total_amount"`
}

type TradeAmountRankResponse struct {
	Sell            []*TradeAmount `json:"sell"`
	Buy             []*TradeAmount `json:"buy"`
	TotalSellUsers  int            `json:"total_sell_users"`
	TotalBuyUsers   int            `json:"total_buy_users"`
	TotalSellAmount string         `json:"total_sell_amount"`
	TotalBuyAmount  string         `json:"total_buy_amount"`
	TotalUsers      int            `json:"total_users"`
}

type TradeDealSummaryRequest struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

type TradeDealSummaryResponse struct {
	DealUsers int `json:"deal_users"`
}

type TradeUsersVolumeRequest struct {
	Market    []string `json:"market"`
	User      []int64  `json:"user"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

type Volume [2]string

type UserVolume map[string]Volume

type TradeUsersVolumeResponse map[string]UserVolume

type tradeClient struct {
	cli matchclient.Invoker
}

func NewTradeClient(c matchclient.Invoker) TradeClient {
	return &tradeClient{cli: c}
}

func (c *tradeClient) TradeNetRank(ctx context.Context, req *TradeNetRankRequest) (reply *TradeNetRankResponse, err error) {
	err = c.cli.Invoke(ctx, "trade.net_rank", &reply, req.Market, req.StartTime, req.EndTime)
	return
}

func (c *tradeClient) TradeAmountRank(ctx context.Context, req *TradeAmountRankRequest) (reply *TradeAmountRankResponse, err error) {
	err = c.cli.Invoke(ctx, "trade.amount_rank", &reply, req.Market, req.StartTime, req.EndTime)
	return
}

func (c *tradeClient) TradeDealSummary(ctx context.Context, req *TradeDealSummaryRequest) (reply *TradeDealSummaryResponse, err error) {
	err = c.cli.Invoke(ctx, "trade.amount_rank", &reply, req.StartTime, req.EndTime)
	return
}

func (c *tradeClient) TradeUsersVolume(ctx context.Context, req *TradeUsersVolumeRequest) (reply TradeUsersVolumeResponse, err error) {
	err = c.cli.Invoke(ctx, "trade.users_volume", &reply, req.Market, req.User, req.StartTime, req.EndTime)
	return
}
