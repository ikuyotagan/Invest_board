package sqlstore

import (
	"database/sql"

	"github.com/Artemchikus/api/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                       *sql.DB
	userRepository           *UserRepository
	stockRepository          *StockRepository
	candelRepository         *CandelRepository
	personalStockRepository *PersonalStockRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Stock() store.StockRepository {
	if s.stockRepository != nil {
		return s.stockRepository
	}

	s.stockRepository = &StockRepository{
		store: s,
	}

	return s.stockRepository
}

func (s *Store) Candel() store.CandelRepository {
	if s.candelRepository != nil {
		return s.candelRepository
	}

	s.candelRepository = &CandelRepository{
		store: s,
	}

	return s.candelRepository
}

func (s *Store) PersonalStock() store.PersonalStockRepository {
	if s.personalStockRepository != nil {
		return s.personalStockRepository
	}

	s.personalStockRepository = &PersonalStockRepository{
		store: s,
	}

	return s.personalStockRepository
}
