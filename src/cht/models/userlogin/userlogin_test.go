package userlogin

import (
	_ "cht/initial"
	"testing"
)

func NewUserlLoginRequest(username string, password string, loginip string, log string) *UserlLoginRequest {
	ulr := UserlLoginRequest{
		Username:             username,
		Password:             password,
		IP:                   loginip,
		ChengHuiTongTraceLog: log,
	}
	return &ulr
}

func TestGetLoginFailedTimes(t *testing.T) {
	req := NewUserlLoginRequest("74965445", "", "", "")
	times, err := GetLoginFailedTimes(req)
	if err != nil {
		t.Fatal("TestGetLoginFailedTimes failed", err)
	}
	t.Log("TestGetLoginFailedTimes ", times)
}

func TestCheckLoginUserExists(t *testing.T) {
	req := NewUserlLoginRequest("111", "", "", "")
	_, _, err := CheckLoginUserExists(req)
	if err != nil {
		t.Fatal("TestCheckLoginUserExists failed", err)
	}
}

func TestCheckpassword(t *testing.T) {
	ulr := NewUserlLoginRequest("july", "9f7add09b41ac15889441e467ff208bf", "", "")
	b := Checkpassword(ulr)
	if b == false {
		t.Fatal("TestCheckpassword failed")
	}
}

func TestInsertUserTimesTb(t *testing.T) {
	ulr := NewUserlLoginRequest("mamaipi", "9f7add09b41ac15889441e467ff208bf", "192.168.8.209", "")
	b, err := InsertUserTimesTb(ulr)
	if err != nil || b == false {
		t.Fatalf("TestInsertUserTimesTb failed", err)
	}
}

func TestCheckUserTimesTbExist(t *testing.T) {
	ulr := NewUserlLoginRequest("mamaipi", "9f7add09b41ac15889441e467ff208bf", "", "")
	b, err := CheckUserTimesTbExist(ulr)
	if b == false || err != nil {
		t.Fatalf("TestCheckUserTimesTbExist failed", err)
	}
}

func TestUpdateUserTimesTb(t *testing.T) {
	ulr := NewUserlLoginRequest("mamaipi", "9f7add09b41ac15889441e467ff208bf", "192.168.8.35", "test.log")
	b, err := UpdateUserTimesTb(ulr)
	if b == false || err != nil {
		t.Fatalf("TestUpdateUserTimesTb failed", err)
	}
}

func TestDeleteUserTimesTb(t *testing.T) {
	ulr := NewUserlLoginRequest("mamaipi", "9f7add09b41ac15889441e467ff208bf", "192.168.8.35", "test.log")
	b, err := DeleteUserTimesTb(ulr)
	if b == false || err != nil {
		t.Fatalf("TestDeleteUserTimesTb failed", err)
	}
}
