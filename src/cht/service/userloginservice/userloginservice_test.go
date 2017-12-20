package userloginservice

import (
	_ "cht/initial"
	"testing"
)

func NewUserLoginRequestStruct(username string, password string, loginip string, log string, type1 int32) *UserLoginRequestStruct {
	return &UserLoginRequestStruct{
		Username:             username,
		Password:             password,
		IP:                   loginip,
		ChengHuiTongTraceLog: log,
		Type:                 type1,
	}
}

func TestGetUserLoginInfo(t *testing.T) {
	ulrs := NewUserLoginRequestStruct("July", "9f7add09b41ac15889441e467ff208bf1", "192.168.8.209", "", 1)
	uls := &UserLoginService{}
	res, _ := uls.GetUserLoginInfo(ulrs)
	if res.Status != VERIFY_PASS {
		t.Fatal("TestGetUserLoginInfo failed")
	}
	t.Logf("TestGetUserLoginInfo response:%v", res)
}
