package teststore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestPersonalStockRepository_Create(t *testing.T) {
	s := teststore.New()

	ps := model.TestPersonalStock1(t)
	assert.NoError(t, s.PersonalStock().Create(ps))
	assert.NotNil(t, ps)
}

func TestPersonalStockRepository_FindByUserIDAndStockID(t *testing.T) {
	s := teststore.New()

	ps1 := model.TestPersonalStock1(t)
	_, err := s.PersonalStock().FindByUserIDAndStockID(ps1.UserID, ps1.StockID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.PersonalStock().Create(ps1)
	ps2, err := s.PersonalStock().FindByUserIDAndStockID(ps1.UserID, ps1.StockID)
	assert.NoError(t, err)
	assert.NotNil(t, ps2)
}

func TestPersonalStockRepository_Find(t *testing.T) {
	s := teststore.New()

	ps1 := model.TestPersonalStock1(t)
	_, err := s.PersonalStock().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.PersonalStock().Create(ps1)
	ps2, err := s.PersonalStock().Find(1)
	assert.NoError(t, err)
	assert.NotNil(t, ps2)
}

func TestPersonalStockRepository_FindStocksByUserID(t *testing.T) {
	s := teststore.New()

	ps1 := model.TestPersonalStock1(t)
	ps2 := model.TestPersonalStock2(t)
	_, err := s.PersonalStock().FindStocksByUserID(ps1.UserID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.PersonalStock().Create(ps1)
	s.PersonalStock().Create(ps2)
	slisePs := make([]*model.PersonalStock, 0)
	slisePs = append(slisePs, ps1)
	slisePs = append(slisePs, ps2)
	ps3, err := s.PersonalStock().FindStocksByUserID(ps1.UserID)
	assert.NoError(t, err)
	assert.NotNil(t, ps3)
	assert.Equal(t, slisePs, ps3)
}
