package sqlstore

import (
	"database/sql"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
)

type StockRepository struct {
	store *Store
}

func (r *StockRepository) Create(s *model.Stock) error {
	return r.store.db.QueryRow("INSERT INTO stocks (name, figi) VALUES($1, $2) RETURNING id",
		s.Name,
		s.FIGI,
	).Scan(&s.ID)
}

func (r *StockRepository) Find(id int) (*model.Stock, error) {
	s := &model.Stock{}

	if err := r.store.db.QueryRow(
		"SELECT id, name, figi FROM stocks WHERE id = $1",
		id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.FIGI,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

func (r *StockRepository) FindByName(name string) (*model.Stock, error) {
	s := &model.Stock{}

	if err := r.store.db.QueryRow(
		"SELECT id, name, figi FROM stocks WHERE name = $1",
		name,
	).Scan(
		&s.ID,
		&s.Name,
		&s.FIGI,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}

func (r *StockRepository) FindByFIGI(figi string) (*model.Stock, error) {
	s := &model.Stock{}
	
	if err := r.store.db.QueryRow(
		"SELECT id, name, figi FROM stocks WHERE figi = $1",
		figi,
	).Scan(
		&s.ID,
		&s.Name,
		&s.FIGI,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return s, nil
}