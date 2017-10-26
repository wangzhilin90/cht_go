package customerlist

import (
	_ "cht/initial"
	"testing"
)

func NewCustomerListRequestStruct(custormerId int32, username string) *CustomerListRequestStruct {
	return &CustomerListRequestStruct{
		Customer: custormerId,
		Username: username,
		Islock:   -1,
	}
}

func TestGetCustomerList(t *testing.T) {
	clr := NewCustomerListRequestStruct(89, "Jerry114")
	cls := customerlistservice{}
	res, _ := cls.GetCustomerList(clr)
	if res.Status != QUERY_CUSTOMER_LIST_SUCCESS {
		t.Fatalf("TestGetCustomerList failed")
	}
	t.Logf("TestGetCustomerList response:%v", res)
}
