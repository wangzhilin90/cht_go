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
	dds := dutydetailsservice{}
	res, err := dds.GetDutyDetails(ddrs)
	if err != nil {
		t.Fatalf("TestGetDutyDetails query failed", err)
	}
	t.Logf("TestGetDutyDetails return value %v", res)
}
