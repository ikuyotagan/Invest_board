package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@gmail.com",
		Password: "password",
	}
}
