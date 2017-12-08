package userappbank

import (
	_ "cht/initial"
	"testing"
)

func NewUserAppBankDetailsRequest(userID int32) *UserAppBankDetailsRequest {
	return &UserAppBankDetailsRequest{
		UserID: userID,
	}
}

func NewUserAppBankUpdateRequest(userID int32) *UserAppBankUpdateRequest {
	return &UserAppBankUpdateRequest{
		UserID:  userID,
		Account: "12345",
	}
}

func NewUserAppBankInsertRequest(userID int32, account string, name string, branch string, addIP string) *UserAppBankInsertRequest {
	return &UserAppBankInsertRequest{
		UserID:  userID,
		Account: account,
		Name:    name,
		Branch:  branch,
		Addip:   addIP,
	}
}

func NewDeletetUserAppBank(userID int32) *UserAppBankDeleteRequest {
	return &UserAppBankDeleteRequest{
		UserID: userID,
	}
}

func TestGetUserAppBankDetails(t *testing.T) {
	uabdr := NewUserAppBankDetailsRequest(3022)
	res, err := GetUserAppBankDetails(uabdr)
	if err != nil {
		t.Fatalf("TestGetUserAppBankDetails failed:%v", err)
	}
	t.Logf("TestGetUserAppBankDetails return value:%v", res)
}

func TestUpdateUserAppBank(t *testing.T) {
	uabur := NewUserAppBankUpdateRequest(29)
	b := UpdateUserAppBank(uabur)
	if b == false {
		t.Fatalf("TestUpdateUserAppBank failed")
	}
}

func TestInsertUserAppBank(t *testing.T) {
	uabir := NewUserAppBankInsertRequest(29, "123434325", "wzl", "深圳支行", "129.19.29.2")
	b := InsertUserAppBank(uabir)
	if b == false {
		t.Fatalf("TestInsertUserAppBank failed")
	}
}

func TestDeletetUserAppBank(t *testing.T) {
	duab := NewDeletetUserAppBank(29)
	b := DeletetUserAppBank(duab)
	if b == false {
		t.Fatalf("TestDeletetUserAppBank failed")
	}
}
