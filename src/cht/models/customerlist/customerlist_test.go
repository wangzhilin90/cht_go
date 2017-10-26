package customerlist

import (
	_ "cht/initial"
	"testing"
)

func NewCustomerListRequest(custormerId int32, username string) *CustomerListRequest {
	return &CustomerListRequest{
		Customer: custormerId,
		Username: username,
		Islock:   -1,
	}
}

func TestGetCustomerListTotalNum(t *testing.T) {
	clr := NewCustomerListRequest(89, "Jerry114")
	num, err := GetCustomerListTotalNum(clr)
	if err != nil {
		t.Fatalf("TestGetCustomerListTotalNum failed:%v", err)
	}
	t.Logf("TestGetCustomerListTotalNum return num:%v", num)
}

func TestGetCustomerList(t *testing.T) {
	clr := NewCustomerListRequest(89, "Jerry114")
	res, err := GetCustomerList(clr)
	if err != nil {
		t.Fatalf("TestGetCustomerList failed:%v", err)
	}
	t.Logf("TestGetCustomerList return value:%v", res)
}
