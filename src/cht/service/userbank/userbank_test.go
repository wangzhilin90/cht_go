package userbank

import (
	_ "cht/initial"
	"testing"
)

func NewUserBankDetailsRequestStruct(user_id int32) *UserBankDetailsRequestStruct {
	return &UserBankDetailsRequestStruct{
		UserID: user_id,
	}
}

func NewUserBankUpdateRequestStruct(user_id int32, name string, account string, city int32) *UserBankUpdateRequestStruct {
	return &UserBankUpdateRequestStruct{
		UserID:  user_id,
		Name:    name,
		Account: account,
		City:    city,
	}
}

func NewUserBankInsertRequestStruct(user_id int32, name string, account string, city int32, branch string, addip string) *UserBankInsertRequestStruct {
	return &UserBankInsertRequestStruct{
		UserID:  user_id,
		Name:    name,
		Account: account,
		City:    city,
		Branch:  branch,
		Addip:   addip,
	}
}

func NewUserBankCountRequestStruct(user_id int32) *UserBankCountRequestStruct {
	return &UserBankCountRequestStruct{
		UserID: user_id,
	}
}

func TestGetUserBankDetails(t *testing.T) {
	ubdr := NewUserBankDetailsRequestStruct(2)
	ubs := userbankservice{}
	res, _ := ubs.GetUserBankDetails(ubdr)
	if res.Status != QUERY_USER_BANK_DETAILS_SUCCESS {
		t.Fatalf("TestGetUserBankDetails query failed")
	}
	t.Logf("TestGetUserBankDetails return value:%v", res)
}

func TestUpdateUserBank(t *testing.T) {
	ubur := NewUserBankUpdateRequestStruct(29, "", "66666", 4)
	ubs := userbankservice{}
	res, _ := ubs.UpdateUserBank(ubur)
	if res.Status != UPDATE_USER_BANK_SUCCESS {
		t.Fatalf("TestUpdateUserBank update failed")
	}
}

func TestInsertUserBank(t *testing.T) {
	ubir := NewUserBankInsertRequestStruct(29, "wzl", "123435", 5, "深圳支行", "192.14.2.2")
	ubs := userbankservice{}
	res, _ := ubs.InsertUserBank(ubir)
	if res.Status != INSERT_USER_BANK_SUCCESS {
		t.Fatalf("TestInsertUserBank insert failed")
	}
}

func TestGetUserBankNum(t *testing.T) {
	ubcr := NewUserBankCountRequestStruct(304324)
	ubs := userbankservice{}
	res, _ := ubs.GetUserBankNum(ubcr)
	if res.Status != COUNT_USER_BANK_SUCCESS {
		t.Fatalf("TestGetUserBankNum query failed")
	}
	t.Logf("TestGetUserBankNum return value:%v", res)
}
