package kefudutydetails

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyDetailsRequestStruct(id int32) *KefuDutyDetailsRequestStruct {
	return &KefuDutyDetailsRequestStruct{
		ID: id,
	}
}

func TestGetKefuDutyDetails(t *testing.T) {
	kddrs := NewKefuDutyDetailsRequestStruct(326)
	kdds := kefudutydetailsservice{}
	res, _ := kdds.GetKefuDutyDetails(kddrs)
	if res.Status != QUERY_KEFU_DUTY_DETAILS_SUCCESS {
		t.Fatalf("TestGetKefuDutyDetails failed")
	}
	t.Logf("TestGetKefuDutyDetails response:%v", res)
}
