package userloginservice

import (
	_ "cht/initial"
	"testing"
)

func TestGetUserLoginInfo(t *testing.T) {
	ulrs := NewUserLoginRequestStruct("July", "9f7add09b41ac15889441e467ff208bf", "", "")
	uls := &UserLoginService{}
	res, err := uls.GetUserLoginInfo(ulrs)
	if err != nil {
		t.Fatal("TestGetUserLoginInfo failed", err)
	}
	t.Logf("TestGetUserLoginInfo response:%v", res)
}
