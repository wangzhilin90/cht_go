package borrowerthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewBorrowerInfoRequestStruct(username string) *BorrowerInfoRequestStruct {
	return &BorrowerInfoRequestStruct{
		Name: username,
	}
}

func TestGetBorrowerInfo(t *testing.T) {
	bir := NewBorrowerInfoRequestStruct("可可")
	bs := &borrowerservice{}
	res, err := bs.GetBorrowerInfo(bir)
	if err != nil {
		t.Fatalf("TestGetBorrowerInfo failed %v", err)
	}
	t.Logf("TestGetBorrowerInfo res %v", res)
}
