package model

type PersonalStock struct {
	ID             int    `json:"-"`
	UserID         int    `json:"-"`
	StockID        int    `json:"-"`
	UserStockValue float32 `json:"value"`
}
