package futuresapi

type Client struct {
	Invoker
	Asset    AssetClient
	Config   ConfigClient
	Index    IndexClient
	Market   MarketClient
	Order    OrderClient
	Position PositionClient
	Trade    TradeClient
	Monitor  MonitorClient
}

func NewClient(invoker Invoker) *Client {
	return &Client{
		Invoker:  invoker,
		Asset:    NewAssetClient(invoker),
		Config:   NewConfigClient(invoker),
		Index:    NewIndexClient(invoker),
		Market:   NewMarketClient(invoker),
		Order:    NewOrderClient(invoker),
		Position: NewPositionClient(invoker),
		Trade:    NewTradeClient(invoker),
		Monitor:  NewMonitorClient(invoker),
	}
}
