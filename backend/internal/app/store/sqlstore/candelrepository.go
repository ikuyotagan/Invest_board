package sqlstore

import (
	"database/sql"
	"time"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type CandelRepository struct {
	store *Store
}

func (r *CandelRepository) Create(c *model.Candel) error {
	return r.store.db.QueryRow("INSERT INTO candels (time, open_price, close_price, lowest_price, highest_price, trading_volume, stock_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		c.Time,
		c.OpenPrice,
		c.ClosePrice,
		c.LowestPrice,
		c.HighestPrice,
		c.TradingVolume,
		c.StockID,
	).Scan(&c.ID)
}

func (r *CandelRepository) Find(id int) (*model.Candel, error) {
	c := &model.Candel{}
	if err := r.store.db.QueryRow(
		"SELECT id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id FROM candels WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.OpenPrice,
		&c.ClosePrice,
		&c.LowestPrice,
		&c.HighestPrice,
		&c.Time,
		&c.TradingVolume,
		&c.StockID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}

func (r *CandelRepository) FindByTimeAndStockID(date time.Time, stockId int) (*model.Candel, error) {
	c := &model.Candel{}
	if err := r.store.db.QueryRow(
		"SELECT id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id FROM candels WHERE (time = $1 AND stock_id = $2)",
		date,
		stockId,
	).Scan(
		&c.ID,
		&c.OpenPrice,
		&c.ClosePrice,
		&c.LowestPrice,
		&c.HighestPrice,
		&c.Time,
		&c.TradingVolume,
		&c.StockID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}

func (r *CandelRepository) FindbyPeriodAndStokID(start, end time.Time, stockId int) ([]*model.Candel, error) {
	c := &model.Candel{}
	arrayC := make([]*model.Candel, 0)

	if err := r.store.db.QueryRow(
		"SELECT id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id FROM candels WHERE ((time >= $1 AND time <= $2) AND stock_id = $3)",
		start,
		end,
		stockId,
	).Scan(
		&c.ID,
		&c.OpenPrice,
		&c.ClosePrice,
		&c.LowestPrice,
		&c.HighestPrice,
		&c.Time,
		&c.TradingVolume,
		&c.StockID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	rows, err := r.store.db.Query(
		"SELECT id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id FROM candels WHERE ((time >= $1 AND time <= $2) AND stock_id = $3)",
		start,
		end,
		stockId,
	)
	if err != nil {
		return nil, store.ErrRecordNotFound
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&c.ID,
			&c.OpenPrice,
			&c.ClosePrice,
			&c.LowestPrice,
			&c.HighestPrice,
			&c.Time,
			&c.TradingVolume,
			&c.StockID,
		); err != nil {
			return nil, err
		}

		candel := &model.Candel{
			ID:            c.ID,
			Time:          c.Time,
			OpenPrice:     c.OpenPrice,
			ClosePrice:    c.ClosePrice,
			HighestPrice:  c.HighestPrice,
			LowestPrice:   c.LowestPrice,
			TradingVolume: c.TradingVolume,
			StockID:       c.StockID,
		}

		arrayC = append(arrayC, candel)
	}

	if err := rows.Err(); err != nil {
		return nil, store.ErrRecordNotFound
	}

	return arrayC, nil
}

func (r *CandelRepository) FindLastByStockID(stockId int) (*model.Candel, error) {
	c := &model.Candel{}
	if err := r.store.db.QueryRow(
		"SELECT id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id FROM candels WHERE stock_id = $1 ORDER BY time DESC LIMIT 1",
		stockId,
	).Scan(
		&c.ID,
		&c.OpenPrice,
		&c.ClosePrice,
		&c.LowestPrice,
		&c.HighestPrice,
		&c.Time,
		&c.TradingVolume,
		&c.StockID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}

func (r *CandelRepository) FindLastByStockFIGI(figi string) (*model.Candel, error) {
	c := &model.Candel{}
	if err := r.store.db.QueryRow(
		"SELECT candels.id, open_price, close_price, lowest_price, highest_price, time, trading_volume, stock_id, figi FROM candels, stocks WHERE figi = $1 and stock_id = stocks.id ORDER BY time DESC LIMIT 1",
		figi,
	).Scan(
		&c.ID,
		&c.OpenPrice,
		&c.ClosePrice,
		&c.LowestPrice,
		&c.HighestPrice,
		&c.Time,
		&c.TradingVolume,
		&c.StockID,
		&c.FIGI,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return c, nil
}
