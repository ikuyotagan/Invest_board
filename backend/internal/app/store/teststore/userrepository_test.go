package teststore_test

import (
	"fmt"
	"testing"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/Artemchikus/api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())


	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u1 := model.TestUser(t)
	_, err := s.User().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().Find(1)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_SetTinkoffKey(t *testing.T) {
	s := teststore.New()

	tu := model.TinkoffAPITestUser(t)
	u := model.TestUser(t)
	s.User().Create(u)
	tu.ID = 1
	assert.NoError(t, s.User().SetTinkoffKey(tu))

	u2, _ := s.User().Find(1)
	fmt.Println(u2)
	assert.Equal(t, tu.TinkoffAPIKey, u2.EncryptedTinkoffAPIKey)
}

func TestUserRepository_IsTinkoffKey(t *testing.T) {
	s := teststore.New()
	
	u := model.TestUser(t)
	s.User().Create(u)
	err := s.User().IsTinkoffKey(1)
	assert.EqualError(t, err, store.ErrNoTinkoffKey.Error())

	tu := model.TinkoffAPITestUser(t)
	tu.ID = 1
	s.User().SetTinkoffKey(tu)
	assert.NoError(t, s.User().IsTinkoffKey(1))
}
