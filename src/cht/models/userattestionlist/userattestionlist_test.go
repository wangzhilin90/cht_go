package userattestionlist

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionListRequest(username, realname string, phone_status, limitnum, offset int32) *UserAttestionListRequest {
	return &UserAttestionListRequest{
		Username:    username,
		Realname:    realname,
		PhoneStatus: phone_status,
		LimitNum:    limitnum,
		LimitOffset: offset,
	}
}

func TestGetUserAttestionTatalNum(t *testing.T) {
	ualr := NewUserAttestionListRequest("kp2pz12u", " 张保华 ", 2, 10, 2)
	num, err := GetUserAttestionTatalNum(ualr)
	if err != nil {
		t.Fatalf("TestGetUserAttestionTatalNum query failed")
	}
	t.Logf("TestGetUserAttestionTatalNum return num:%v", num)
}

func TestGetUserAttestionList(t *testing.T) {
	ualr := NewUserAttestionListRequest("", " ", 2, 11, 0)
	res, err := GetUserAttestionList(ualr)
	if err != nil {
		t.Fatalf("TestGetUserAttestionList failed:%v", err)
	}
	t.Logf("TestGetUserAttestionList res:%v", res)
}
