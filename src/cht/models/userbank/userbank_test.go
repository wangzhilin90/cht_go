package userbank

import (
	_ "cht/initial"
	"testing"
)

func NewUserBankDetailsRequest(user_id int32) *UserBankDetailsRequest {
	return &UserBankDetailsRequest{
		UserID: user_id,
	}
}

func NewUserBankUpdateRequest(user_id int32, name string, account string, city int32) *UserBankUpdateRequest {
	return &UserBankUpdateRequest{
		UserID:  user_id,
		Name:    name,
		Account: account,
		City:    city,
	}
}

func NewUserBankInsertRequest(user_id int32, name string, account string, city int32, branch string, addip string) *UserBankInsertRequest {
	return &UserBankInsertRequest{
		UserID:  user_id,
		Name:    name,
		Account: account,
		City:    city,
		Branch:  branch,
		Addip:   addip,
	}
}

func NewUserBankCountRequest(user_id int32) *UserBankCountRequest {
	return &UserBankCountRequest{
		UserID: user_id,
	}
}

func TestGetUserBankDetails(t *testing.T) {
	ubdr := NewUserBankDetailsRequest(2)
	res, err := GetUserBankDetails(ubdr)
	if err != nil {
		t.Fatalf("TestGetUserBankDetails query failed:%v", err)
	}
	t.Logf("TestGetUserBankDetails return value:%v", res)
}

func TestUpdateUserBank(t *testing.T) {
	ubur := NewUserBankUpdateRequest(29, "", "55555", 4)
	b := UpdateUserBank(ubur)
	if b == false {
		t.Fatalf("TestUpdateUserBank update failed")
	}
}

func TestInsertUserBank(t *testing.T) {
	ubir := NewUserBankInsertRequest(29, "wzl", "123435", 5, "深圳支行", "192.14.2.2")
	b := InsertUserBank(ubir)
	if b == false {
		t.Fatalf("TestInsertUserBank insert failed")
	}
}

func TestGetUserBankNum(t *testing.T) {
	ubcr := NewUserBankCountRequest(304324)
	num, err := GetUserBankNum(ubcr)
	if err != nil {
		t.Fatalf("TestGetUserBankNum query failed:%v", err)
	}
	t.Logf("TestGetUserBankNum return value:%v", num)
}
