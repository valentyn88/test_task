package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	errMsgParams  = "not enough params to create new user"
	errMsgConvert = "could not convert row id to int64: %s %s"
)

// ConvStr2User - convert string to User
func ConvStr2User(s string) (*User, error) {
	row := strings.Split(s, ",")
	if len(row) < 4 {
		return nil, errors.New(errMsgParams)
	}

	id, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf(errMsgConvert, row[0], err)
	}

	u := &User{
		Id:           id,
		Name:         row[1],
		Email:        row[2],
		MobileNumber: row[3],
	}

	return u, nil
}
