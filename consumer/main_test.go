package main

import (
	"context"
	"github.com/matryer/is"
	"test_task/user"
	"testing"
)

type mockStorage struct {
}

func (ms mockStorage) AddRow(u user.User) user.User {
	return user.User{Id: 1, Name: "test1", Email: "test1@example.com", MobileNumber: "12345"}
}

func TestUserServer_Add(t *testing.T) {
	us := userServer{mockStorage{}}
	ist := is.New(t)
	expectedUsr := &user.User{Id: 1, Name: "test1", Email: "test1@example.com", MobileNumber: "12345"}
	usr, err := us.Add(context.Background(), &user.User{})
	ist.NoErr(err)              // an error occured while ading user
	ist.Equal(expectedUsr, usr) // expected and got results are unequal
}
