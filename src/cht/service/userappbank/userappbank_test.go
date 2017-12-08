package userappbank

import (
	_ "cht/initial"
	"testing"
)

func NewUserAppBankDetailsRequestStruct(userID int32) *UserAppBankDetailsRequestStruct {
	return &UserAppBankDetailsRequestStruct{
		UserID: userID,
	}
}

func NewUserAppBankUpdateRequestStruct(userID int32) *UserAppBankUpdateRequestStruct {
	return &UserAppBankUpdateRequestStruct{
		UserID:  userID,
		Account: "12345",
	}
}

func NewUserAppBankInsertRequestStruct(userID int32, account string, name string, branch string, addIP string) *UserAppBankInsertRequestStruct {
	return &UserAppBankInsertRequestStruct{
		UserID:  userID,
		Account: account,
		Name:    name,
		Branch:  branch,
		Addip:   addIP,
	}
}

func NewDeletetUserAppBankStruct(userID int32) *UserAppBankDeleteRequestStruct {
	return &UserAppBankDeleteRequestStruct{
		UserID: userID,
	}
}

func TestGetUserAppBankDetails(t *testing.T) {
	uabdr := NewUserAppBankDetailsRequestStruct(29)
	uabs := userappbankservice{}
	res, _ := uabs.GetUserAppBankDetails(uabdr)
	if res.Status != QUERY_USER_BANK_DETAILS_SUCCESS {
		t.Fatalf("TestGetUserAppBankDetails failed")
	}
	t.Logf("TestGetUserAppBankDetails return value:%v", res)
}

func TestUpdateUserAppBank(t *testing.T) {
	uabur := NewUserAppBankUpdateRequestStruct(29)
	uabs := userappbankservice{}
	res, _ := uabs.UpdateUserAppBank(uabur)
	if res.Status != UPDATE_USER_APP_BANK_SUCCESS {
		t.Fatalf("TestUpdateUserAppBank failed")
	}
}

func TestInsertUserAppBank(t *testing.T) {
	uabir := NewUserAppBankInsertRequestStruct(29, "123434325", "wzl", "深圳支行", "129.19.29.2")
	uabs := userappbankservice{}
	res, _ := uabs.InsertUserAppBank(uabir)
	if res.Status != INSERT_USER_APP_BANK_SUCCESS {
		t.Fatalf("TestInsertUserAppBank failed")
	}
}

func TestDeletetUserAppBank(t *testing.T) {
	duab := NewDeletetUserAppBankStruct(29)
	uabs := userappbankservice{}
	res, _ := uabs.DeletetUserAppBank(duab)
	if res.Status != DELETE_USER_APP_BANK_SUCCESS {
		t.Fatalf("TestDeletetUserAppBank failed")
	}
}
