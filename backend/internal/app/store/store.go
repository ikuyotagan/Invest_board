package store

type Store interface {
	User() UserRepository
	Stock() StockRepository
	Candel() CandelRepository
	PersonalStock() PersonalStockRepository
}
