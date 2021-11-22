package teststore

import (
	"time"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type CandelRepository struct {
	store   *Store
	candels map[int]*model.Candel
}

func (r *CandelRepository) Create(c *model.Candel) error {
	c.ID = len(r.candels) + 1
	r.candels[c.ID] = c

	return nil
}

func (r *CandelRepository) Find(id int) (*model.Candel, error) {
	for _, c := range r.candels {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *CandelRepository) FindByTimeAndStockID(date time.Time, stockId int) (*model.Candel, error) {
	for _, c := range r.candels {
		if c.StockID == stockId && c.Time == date {
			return c, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *CandelRepository) FindbyPeriodAndStokID(start, end time.Time, stockId int) ([]*model.Candel, error) {
	arrayC := make([]*model.Candel, 0)

	for _, c := range r.candels {
		if (c.Time.After(start) || c.Time == start) && (c.Time.Before(end) || c.Time == end) && c.StockID == stockId {
			arrayC = append(arrayC, c)
		}
	}

	if len(arrayC) == 0 {
		return nil, store.ErrRecordNotFound
	}

	return arrayC, nil
}

func (r *CandelRepository) FindLastByStockID(stockId int) (*model.Candel, error) {
	max, _ := time.Parse(time.RFC3339Nano, "2000-01-01T15:35:00Z")
	start := max
	maxCandel := &model.Candel{}

	for _, c := range r.candels {
		if c.StockID == stockId && c.Time.After(max) {
			max = c.Time
			maxCandel = c
		}
	}

	if max.Equal(start) {
		return nil, store.ErrRecordNotFound
	}
	
	return maxCandel, nil
}
