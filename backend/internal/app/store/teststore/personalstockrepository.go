package teststore

import (
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type PersonalStockRepository struct {
	store          *Store
	personalStocks map[int]*model.PersonalStock
}

func (r *PersonalStockRepository) Create(ps *model.PersonalStock) error {
	ps.ID = len(r.personalStocks) + 1
	r.personalStocks[ps.ID] = ps

	return nil
}

func (r *PersonalStockRepository) Find(id int) (*model.PersonalStock, error) {
	for _, ps := range r.personalStocks {
		if ps.ID == id {
			return ps, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *PersonalStockRepository) FindStocksByUserID(userId int) ([]*model.PersonalStock, error) {
	arrayPs := make([]*model.PersonalStock, 0)

	for _, ps := range r.personalStocks {
		if ps.UserID == userId {
			arrayPs = append(arrayPs, ps)
		}
	}

	if len(arrayPs) == 0 {
		return nil, store.ErrRecordNotFound
	}

	return arrayPs, nil
}

func (r *PersonalStockRepository) FindByUserIDAndStockID(userId, stockId int) (*model.PersonalStock, error) {
	for _, ps := range r.personalStocks {
		if ps.UserID == userId && ps.StockID == stockId {
			return ps, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *PersonalStockRepository) UpdateBalance(ps *model.PersonalStock) error {
	for _, personalStocks := range r.personalStocks {
		if personalStocks.UserID == ps.UserID && personalStocks.StockID == ps.StockID {
			personalStocks.UserStockValue = ps.UserStockValue
			return nil
		}
	}
	return store.ErrRecordNotFound
}
