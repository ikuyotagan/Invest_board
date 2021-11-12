package teststore

import (
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type Store struct {
	userRepository          *UserRepository
	stockRepository         *StockRepository
	candelRepository        *CandelRepository
	personalStockRepository *PersonalStockRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Stock() store.StockRepository {
	if s.stockRepository != nil {
		return s.stockRepository
	}

	s.stockRepository = &StockRepository{
		store:  s,
		stocks: make(map[int]*model.Stock),
	}

	return s.stockRepository
}

func (s *Store) Candel() store.CandelRepository {
	if s.candelRepository != nil {
		return s.candelRepository
	}

	s.candelRepository = &CandelRepository{
		store:   s,
		candels: make(map[int]*model.Candel),
	}

	return s.candelRepository
}

func (s *Store) PersonalStock() store.PersonalStockRepository {
	if s.personalStockRepository != nil {
		return s.personalStockRepository
	}

	s.personalStockRepository = &PersonalStockRepository{
		store:          s,
		personalStocks: make(map[int]*model.PersonalStock),
	}

	return s.personalStockRepository
}
