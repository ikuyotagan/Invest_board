package sqlstore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestCandelRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("candels")

	s := sqlstore.New(db)
	c := model.TestCandel1(t)
	assert.NoError(t, s.Candel().Create(c))
	assert.NotNil(t, c)
}

func TestCandelRepository_FindByTimeAndStockID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("candels")

	s := sqlstore.New(db)
	c1 := model.TestCandel1(t)
	_, err := s.Candel().FindByTimeAndStockID(c1.Time, c1.StockID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	
	s.Candel().Create(c1)
	c2, err := s.Candel().FindByTimeAndStockID(c1.Time, c1.StockID)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
}

func TestCandelRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("candels")

	s := sqlstore.New(db)
	c1 := model.TestCandel1(t)
	_, err := s.Candel().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	
	s.Candel().Create(c1)
	c2, err := s.Candel().Find(c1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
}

func TestCandelRepository_FindbyPeriodAndStokID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("candels")

	s := sqlstore.New(db)
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