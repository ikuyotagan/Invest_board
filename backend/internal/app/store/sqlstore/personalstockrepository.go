package sqlstore

import (
	"database/sql"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type PersonalStockRepository struct {
	store *Store
}

func (r *PersonalStockRepository) Create(ps *model.PersonalStock) error {
	return r.store.db.QueryRow("INSERT INTO personal_stocks (user_id, stock_id, user_stock_value) VALUES($1, $2, $3) RETURNING id",
		ps.UserID,
		ps.StockID,
		ps.UserStockValue,
	).Scan(&ps.ID)
}

func (r *PersonalStockRepository) UpdateBalance(ps *model.PersonalStock) error {
	return r.store.db.QueryRow("UPDATE personal_stocks SET user_stock_value=$1 WHERE user_id=$2 AND stock_id=$3 RETURNING id",
		ps.UserStockValue,
		ps.UserID,
		ps.StockID,
	).Scan(&ps.ID)
}

func (r *PersonalStockRepository) Find(id int) (*model.PersonalStock, error) {
	ps := &model.PersonalStock{}
	if err := r.store.db.QueryRow(
		"SELECT id, stock_id, user_id, user_stock_value FROM personal_stocks WHERE id = $1",
		id,
	).Scan(
		&ps.ID,
		&ps.StockID,
		&ps.UserID,
		&ps.UserStockValue,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return ps, nil
}

func (r *PersonalStockRepository) FindStocksByUserID(userId int) ([]*model.PersonalStock, error) {
	ps := &model.PersonalStock{}
	arrayPs := make([]*model.PersonalStock, 0)

	if err := r.store.db.QueryRow(
		"SELECT id, stock_id, user_id, user_stock_value FROM personal_stocks WHERE user_id = $1",
		userId,
	).Scan(
		&ps.ID,
		&ps.StockID,
		&ps.UserID,
		&ps.UserStockValue,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	rows, err := r.store.db.Query(
		"SELECT personal_stocks.id, stock_id, name, figi, user_id, user_stock_value FROM personal_stocks, stocks WHERE user_id = $1 and stock_id = stocks.id",
		userId,
	)
	if err != nil {
		return nil, store.ErrRecordNotFound
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&ps.ID,
			&ps.StockID,
			&ps.StockName,
			&ps.StockFIGI,
			&ps.UserID,
			&ps.UserStockValue,
		); err != nil {
			return nil, err
		}

		personalStock := &model.PersonalStock{
			ID:             ps.ID,
			StockID:        ps.StockID,
			StockName:      ps.StockName,
			StockFIGI:      ps.StockFIGI,
			UserID:         ps.UserID,
			UserStockValue: ps.UserStockValue,
		}

		arrayPs = append(arrayPs, personalStock)
	}

	if err = rows.Err(); err != nil {
		return nil, store.ErrRecordNotFound
	}

	return arrayPs, nil
}

func (r *PersonalStockRepository) FindByUserIDAndStockID(userId, stockId int) (*model.PersonalStock, error) {
	ps := &model.PersonalStock{}
	if err := r.store.db.QueryRow(
		"SELECT id, stock_id, user_id, user_stock_value FROM personal_stocks WHERE (stock_id = $1 AND user_id = $2)",
		stockId,
		userId,
	).Scan(
		&ps.ID,
		&ps.StockID,
		&ps.UserID,
		&ps.UserStockValue,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return ps, nil
}
