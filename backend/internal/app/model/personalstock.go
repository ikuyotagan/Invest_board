package model

type PersonalStock struct {
	ID             int     `json:"id"`
	UserID         int     `json:"user_id"`
	StockID        int     `json:"stock_id"`
	StockName      string  `json:"stock_name"`
	StockFIGI      string  `json:"stock_figi"`
	UserStockValue float32 `json:"value"`
}
