package borrowuserdetails

import (
	_ "cht/initial"
	"testing"
)

func NewBorrowerInfoRequestStruct(username string) *BorrowUserDetailsRequestStruct {
	return &BorrowUserDetailsRequestStruct{
		Name: username,
	}
}

func TestGetBorrowUserDetails(t *testing.T) {
	bir := NewBorrowerInfoRequestStruct("xiezhenyuan")
	bs := &borrowerservice{}
	res, err := bs.GetBorrowUserDetails(bir)
	if err != nil {
		t.Fatalf("TestGetBorrowerInfo failed %v", err)
	}
	t.Logf("TestGetBorrowerInfo res %v", res)
}
