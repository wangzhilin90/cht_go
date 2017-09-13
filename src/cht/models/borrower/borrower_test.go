package borrower

import (
	_ "cht/initial"
	"testing"
)

func NewBorrowerInfoRequest(username string) *BorrowerInfoRequest {
	return &BorrowerInfoRequest{
		Name: username,
	}
}

func TestGetBorrowerUID(t *testing.T) {
	birs := NewBorrowerInfoRequest("Jerry114")
	res, err := getBorrowerUID(birs)
	if err != nil {
		t.Fatalf("TestGetBorrowerUID failed %v", err)
	}
	t.Logf("TestGetBorrowerUID res:%v", res)
}

func TestGetCardID(t *testing.T) {
	res, err := getCardID(2)
	if err != nil {
		t.Fatalf("TestGetCardID failed %v", err)
	}
	t.Logf("TestGetCardID res:%v", res)
}

func TestGetCreditUse(t *testing.T) {
	res, err := getCreditUse(29)
	if err != nil {
		t.Fatalf("TestGetCreditUse failed %v", err)
	}
	t.Logf("TestGetCreditUse res:%v", res)
}

func TestGetGuarantor(t *testing.T) {
	res, err := getGuarantor(29)
	if err != nil {
		t.Fatalf("TestGetGuarantor failed %v", err)
	}
	t.Logf("TestGetGuarantor res:%v", res)
}

func TestGetMaterialInfo(t *testing.T) {
	res, err := getMaterialInfo(29)
	if err != nil {
		t.Fatalf("TestGetMaterialInfo failed %v", err)
	}
	t.Logf("TestGetMaterialInfo res:%v", res)
}

func TestGetBorrowerInfo(t *testing.T) {
	birs := NewBorrowerInfoRequest("可可")
	res, err := GetBorrowerInfo(birs)
	if err != nil {
		t.Fatalf("TestGetBorrowerInfo failed %v", err)
	}
	t.Logf("TestGetBorrowerInfo res:%v %v", res, res.MaterialInfo)
}
