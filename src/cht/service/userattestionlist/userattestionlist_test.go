package userattestionlist

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionListRequestStruct(username, realname string, phone_status, limitnum, offset int32) *UserAttestionListRequestStruct {
	return &UserAttestionListRequestStruct{
		Username:    username,
		Realname:    realname,
		PhoneStatus: phone_status,
		LimitNum:    limitnum,
		LimitOffset: offset,
	}
}

func TestGetUserAttestionList(t *testing.T) {
	ualr := NewUserAttestionListRequestStruct("", " ", 2, 11, 0)
	uals := userattestionlistservice{}
	res, err := uals.UserAttestionList(ualr)
	if err != nil {
		t.Fatalf("TestGetUserAttestionList failed:%v", err)
	}
	t.Logf("TestGetUserAttestionList res:%v", res)
}
