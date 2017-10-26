package customerupdate

import (
	_ "cht/initial"
	"testing"
)

func NewCustomerUpdateRequestStruct(id string, isLock int32) *CustomerUpdateRequestStruct {
	return &CustomerUpdateRequestStruct{
		ID:     id,
		Islock: isLock,
	}
}

func TestUpdateCustomer(t *testing.T) {
	curs := NewCustomerUpdateRequestStruct("2192,2193,324234,2194", 0)
	cus := customerupdateservice{}
	res, _ := cus.UpdateCustomer(curs)
	if res.Status != UPDATE_CUSTOMER_LOCK_SUCCESS {
		t.Fatalf("TestUpdateCustomer failed")
	}
	t.Logf("TestUpdateCustomer response:%v", res)
}
