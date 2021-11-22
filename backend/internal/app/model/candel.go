package model

import "time"

type Candel struct {
	ID            int       `json:"id"`
	Time          time.Time `json:"time"`
	OpenPrice     float32   `json:"open_price"`
	ClosePrice    float32   `json:"close_price"`
	HighestPrice  float32   `json:"highest_price"`
	LowestPrice   float32   `json:"lowest_price"`
	TradingVolume float32   `json:"volume"`
	StockID       int       `json:"stock_id"`
}
