package model

import "time"

type Coeff struct {
	Name    string // P/E or P/E or P/BV
	Value   float32
	Time    time.Time
	StockId int
}
