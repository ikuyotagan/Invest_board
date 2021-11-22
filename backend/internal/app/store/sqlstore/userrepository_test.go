package sqlstore_test

import (
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_SetTinkoffKey(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	tu := model.TinkoffAPITestUser(t)
	u := model.TestUser(t)
	s.User().Create(u)
	tu.ID = u.ID
	assert.NoError(t, s.User().SetTinkoffKey(tu))

	u2, _ := s.User().Find(u.ID)
	assert.NotNil(t, u2.TinkoffAPIKey)
	assert.Equal(t, u2.TinkoffAPIKey, tu.TinkoffAPIKey)
}

func TestUserRepository_IsTinkoffKey(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	err := s.User().IsTinkoffKey(u.ID)
	assert.EqualError(t, err, store.ErrNoTinkoffKey.Error())

	tu := model.TinkoffAPITestUser(t)
	tu.ID = u.ID
	s.User().SetTinkoffKey(tu)
	assert.NoError(t, s.User().IsTinkoffKey(tu.ID))
}
