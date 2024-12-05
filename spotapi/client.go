package spotapi

import "github.com/ashexchange/matchclient/v2"

type Client struct {
	Asset  AssetClient
	Config ConfigClient
	Market MarketClient
	Order  OrderClient
	Trade  TradeClient
}

func NewClient(invoker matchclient.Invoker) *Client {
	return &Client{
		Asset:  NewAssetClient(invoker),
		Config: NewConfigClient(invoker),
		Market: NewMarketClient(invoker),
		Order:  NewOrderClient(invoker),
		Trade:  NewTradeClient(invoker),
	}
}
