package futuresapi

import (
	"context"
	"fmt"

	"github.com/ashexchange/matchclient/v2/types"
)

type TradeClient struct {
	invoker Invoker
}

func NewTradeClient(invoker Invoker) TradeClient {
	return TradeClient{invoker}
}

type TradeNetRankRequest struct {
	Markets []string
	// StartTime must greater than zero
	StartTime types.Timestamp
	// EndTime must greater than zero
	EndTime types.Timestamp
}

type TradeNetRankResult struct {
	Buy            []*TradeNetRank `json:"buy"`
	TotalBuyUsers  int             `json:"total_buy_users"`
	Sell           []*TradeNetRank `json:"sell"`
	TotalNet       types.Number    `json:"total_net"`
	TotalSellUsers int             `json:"total_sell_users"`
	TotalAmount    types.Number    `json:"total_amount"`
}

type TradeNetRank struct {
	UserId types.UserID `json:"user_id"`
	Total  types.Number `json:"total"`
	Net    types.Number `json:"net"`
}

func (c TradeClient) NetRank(ctx context.Context, req TradeNetRankRequest) (result *TradeNetRankResult, err error) {
	err = c.invoker.Invoke(ctx, "trade.net_rank", &result,
		req.Markets,
		req.StartTime,
		req.EndTime,
	)

	return
}

type TradeAmountRankRequest struct {
	Markets []string
	// StartTime must greater than zero
	StartTime types.Timestamp
	// EndTime must greater than zero
	EndTime types.Timestamp
}

type TradeAmountRankResult struct {
	TotalBuyAmount  types.Number       `json:"total_buy_amount"`
	Buy             []*TradeAmountRank `json:"buy"`
	TotalBuyUsers   int                `json:"total_buy_users"`
	Sell            []*TradeAmountRank `json:"sell"`
	TotalSellUsers  int                `json:"total_sell_users"`
	TotalSellAmount types.Number       `json:"total_sell_amount"`
	TotalUsers      int                `json:"total_users"`
}

type TradeAmountRank struct {
	UserId      types.UserID `json:"user_id"`
	Amount      types.Number `json:"amount"`
	TotalAmount types.Number `json:"total_amount"`
}

func (c TradeClient) AmountRank(ctx context.Context, req TradeAmountRankRequest) (result *TradeAmountRankResult, err error) {
	err = c.invoker.Invoke(ctx, "trade.amount_rank", &result,
		req.Markets,
		req.StartTime,
		req.EndTime,
	)

	return
}

type TradeDealSummaryRequest struct {
	// StartTime must greater than zero
	StartTime types.Timestamp
	// EndTime must greater than zero
	EndTime types.Timestamp
}

type TradeDealSummaryResult struct {
	DealUsers int `json:"deal_users"`
}

func (c TradeClient) DealSummary(ctx context.Context, req TradeDealSummaryRequest) (result *TradeDealSummaryResult, err error) {
	err = c.invoker.Invoke(ctx, "trade.deal_summary", &result,
		req.StartTime,
		req.EndTime,
	)

	return
}

type TradeUsersVolumeRequest struct {
	Markets []string
	Users   []types.UserID
	// StartTime must greater than zero
	StartTime types.Timestamp
	// EndTime must greater than zero
	EndTime types.Timestamp
}

type TradeUserVolumeResult map[string]map[types.UserID]TradeUserStatics

type TradeUserStatics [4]types.Number

func (v TradeUserStatics) String() string {
	return fmt.Sprintf("TradeUserAmount{BuyAmount:%s, SellAmount:%s, MakerVolume:%s, TakerVolume:%s,}",
		v.BuyAmount(),
		v.SellAmount(),
		v.MakerAmount(),
		v.TakerAmount(),
	)
}

func (v TradeUserStatics) BuyAmount() types.Number   { return v[0] }
func (v TradeUserStatics) SellAmount() types.Number  { return v[1] }
func (v TradeUserStatics) MakerAmount() types.Number { return v[2] }
func (v TradeUserStatics) TakerAmount() types.Number { return v[3] }

func (c TradeClient) UsersVolume(ctx context.Context, req TradeUsersVolumeRequest) (result TradeUserVolumeResult, err error) {
	err = c.invoker.Invoke(ctx, "trade.users_volume", &result,
		req.Markets,
		req.Users,
		req.StartTime,
		req.EndTime,
	)

	return
}
