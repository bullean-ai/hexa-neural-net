package binance

type ChartData struct {
	OpenTime         int64   `json:"open_time"`
	OpenPrice        float64 `json:"open_price"`
	HighPrice        float64 `json:"high_price"`
	LowPrice         float64 `json:"low_price"`
	ClosePrice       float64 `json:"close_price"`
	Volume           float64 `json:"volume"`
	CloseTime        int64   `json:"close_time"`
	QuoteAssetVolume float64 `json:"quote_asset_volume,omitempty"`
	AssetVolume      float64 `json:"asset_volume,omitempty"`
	NumberOfTrades   int64   `json:"number_of_trades,omitempty"`
	TakerBaseVolume  float64 `json:"taker_base_volume,omitempty"`
	TakerQuoteVolume float64 `json:"taker_quote_volume,omitempty"`
}

type TakerLongShortRatioData struct {
	BuySellRatio string `json:"buySellRatio"`
	SellVol      string `json:"sellVol"`
	BuyVol       string `json:"buyVol"`
	Timestamp    int64  `json:"timestamp"`
}

type DepthData struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	E            int64      `json:"E"`
	T            int64      `json:"T"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type PriceTickerData struct {
	Pair               string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice          string `json:"lastPrice"`
}
