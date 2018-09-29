package user

import (
	"errors"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestNewUser(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  *User
		err  error
	}{
		{
			"fail - not enough params",
			"1,test,",
			nil,
			errors.New(errMsgParams),
		},
		{
			"fail - could not convert id to int64",
			"test1,test2,test3,test4",
			nil,
			fmt.Errorf(errMsgConvert, "test1", `strconv.ParseInt: parsing "test1": invalid syntax`),
		},
		{
			"success",
			"1,test2,test3,0456",
			&User{Id: 1, Name: "test2", Email: "test3", MobileNumber: "0456"},
			nil,
		},
	}

	ist := is.New(t)

	for _, tc := range testCases {
		t.Log("test case name: ", tc.name)

		u, err := ConvStr2User(tc.in)
		if err != nil {
			ist.Equal(tc.err, err) // expected and got results are unequal
			continue
		}

		ist.Equal(tc.out, u) // expected and got results are unequal
	}
}
