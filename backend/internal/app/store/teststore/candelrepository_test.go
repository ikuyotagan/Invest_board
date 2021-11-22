package teststore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCandelRepository_Create(t *testing.T) {
	s := teststore.New()

	c := model.TestCandel1(t)
	assert.NoError(t, s.Candel().Create(c))
	assert.NotNil(t, c)
}

func TestCandelRepository_FindByTimeAndStockID(t *testing.T) {
	s := teststore.New()

	c1 := model.TestCandel1(t)
	_, err := s.Candel().FindByTimeAndStockID(c1.Time, c1.StockID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	
	s.Candel().Create(c1)
	c2, err := s.Candel().FindByTimeAndStockID(c1.Time, c1.StockID)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
}

func TestCandelRepository_Find(t *testing.T) {
	s := teststore.New()

	c1 := model.TestCandel1(t)
	_, err := s.Candel().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	
	s.Candel().Create(c1)
	c2, err := s.Candel().Find(1)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
}

func TestCandelRepository_FindbyPeriodAndStokID(t *testing.T) {
	s := teststore.New()
	
	c1 := model.TestCandel1(t)
	c2 := model.TestCandel2(t)
	_, err := s.Candel().FindbyPeriodAndStokID(c1.Time, c2.Time, c1.StockID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	
	s.Candel().Create(c1)
	s.Candel().Create(c2)
	sliseC := make([]*model.Candel, 0)
	sliseC = append(sliseC, c1)
	sliseC = append(sliseC, c2)
	c3, err := s.Candel().FindbyPeriodAndStokID(c1.Time, c2.Time, c1.StockID)
	assert.NoError(t, err)
	assert.NotNil(t, c3)
	assert.Equal(t, sliseC, c3)
}

func TestCandelRepository_FindLastByStockID(t *testing.T) {
	s := teststore.New()
	
	c1 := model.TestCandel1(t)
	c2 := model.TestCandel2(t)
	_, err := s.Candel().FindLastByStockID(c1.StockID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Candel().Create(c1)
	s.Candel().Create(c2)
	c3, err := s.Candel().FindLastByStockID(c1.StockID)
	assert.NoError(t, err)
	assert.NotNil(t, c3)
	assert.Equal(t, c2, c3)
}
