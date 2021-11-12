package teststore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestStockRepository_Create(t *testing.T) {
	s := teststore.New()

	st := model.TestStock(t)
	assert.NoError(t, s.Stock().Create(st))
	assert.NotNil(t, st)
}

func TestStockRepository_FindByName(t *testing.T) {
	s := teststore.New()

	st1 := model.TestStock(t)
	_, err := s.Stock().FindByName(st1.Name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().FindByName(st1.Name)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}

func TestStockRepository_Find(t *testing.T) {
	s := teststore.New()

	st1 := model.TestStock(t)
	_, err := s.Stock().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().Find(1)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}

func TestStockRepository_FindByFIGI(t *testing.T) {
	s := teststore.New()
	
	st1 := model.TestStock(t)
	_, err := s.Stock().FindByFIGI(st1.FIGI)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Stock().Create(st1)
	st2, err := s.Stock().FindByFIGI(st1.FIGI)
	assert.NoError(t, err)
	assert.NotNil(t, st2)
}