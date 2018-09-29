package storage

import (
	"test_task/user"
	"testing"

	"github.com/matryer/is"
)

func TestMapStorage_AddRow(t *testing.T) {
	strg := New()

	users := []user.User{
		{Id: 1, Name: "test1", Email: "test1@example.com", MobileNumber: "12345"},
		{Id: 2, Name: "test1", Email: "test2@example.com", MobileNumber: "12345"},
		{Id: 3, Name: "test1", Email: "test3@example.com", MobileNumber: "12345"},
	}

	expected := []int64{1, 2, 3}
	got := make([]int64, len(expected))

	for i, u := range users {
		usr := strg.AddRow(u)
		got[i] = usr.Id
	}

	ist := is.New(t)
	ist.Equal(expected, got) // expected and got results are unequal
}
