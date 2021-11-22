package teststore

import (
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type StockRepository struct {
	store  *Store
	stocks map[int]*model.Stock
}

func (r *StockRepository) Create(s *model.Stock) error {
	s.ID = len(r.stocks) + 1
	r.stocks[s.ID] = s

	return nil
}

func (r *StockRepository) Find(id int) (*model.Stock, error) {
	for _, s := range r.stocks {
		if s.ID == id {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *StockRepository) FindByName(name string) (*model.Stock, error) {
	for _, s := range r.stocks {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *StockRepository) FindByFIGI(figi string) (*model.Stock, error) {
	for _, s := range r.stocks {
		if s.FIGI == figi {
			return s, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *StockRepository) GetAll() ([]*model.Stock, error) {
	arrayS := make([]*model.Stock, 0)

	for _, s := range r.stocks {
		arrayS = append(arrayS, s)
	}

	return arrayS, nil
}
