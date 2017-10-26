package customerupdate

import (
	_ "cht/initial"
	"testing"
)

func NewCustomerUpdateRequest(id string, isLock int32) *CustomerUpdateRequest {
	return &CustomerUpdateRequest{
		ID:     id,
		Islock: isLock,
	}
}

func TestUpdateCustomer(t *testing.T) {
	cur := NewCustomerUpdateRequest("2192,2193,324234,2194", 1)
	b := UpdateCustomer(cur)
	if b == false {
		t.Fatalf("TestUpdateCustomer failed")
	}
}
