package model

import "time"

type Candel struct {
	ID            int    `json:"-"`
	Time          time.Time   `json:"time"`
	OpenPrice     float32 `json:"o"`
	ClosePrice    float32 `json:"c"`
	HighestPrice  float32 `json:"h"`
	LowestPrice   float32 `json:"l"`
	TradingVolume float32 `json:"v"`
	StockID       int    `json:"-"`
}
