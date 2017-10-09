package dutydetails

import (
	_ "cht/initial"
	"testing"
)

func NewDutyDetailsRequestStruct() *DutyDetailsRequestStruct {
	return &DutyDetailsRequestStruct{}
}

func TestGetDutyDetails(t *testing.T) {
	ddrs := NewDutyDetailsRequestStruct()
	res, err := GetDutyDetails(ddrs)
	if err != nil {
		t.Fatalf("TestGetDutyDetails failed :%v", err)
	}
	t.Logf("TestGetDutyDetails res %v", res)
}
