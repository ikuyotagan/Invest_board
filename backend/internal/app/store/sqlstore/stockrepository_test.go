package sqlstore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestStockRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("stocks")

	s := sqlstore.New(db)
	st := model.TestStock(t)
	assert.NoError(t, s.Stock().Create(st))
	assert.NotNil(t, st)
}

func TestStockRepository_FindByName(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("stocks")

	s := sqlstore.New(db)
	st1 := model.TestStock(t)
	_, err := s.Stock().FindByName(st1.Name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().FindByName(st1.Name)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}

func TestStockRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("stocks")

	s := sqlstore.New(db)
	st1 := model.TestStock(t)
	_, err := s.Stock().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().Find(st1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}

func TestStockRepository_FindByFIGI(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("stocks")

	s := sqlstore.New(db)
	st1 := model.TestStock(t)
	_, err := s.Stock().FindByFIGI(st1.FIGI)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().FindByFIGI(st1.FIGI)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}