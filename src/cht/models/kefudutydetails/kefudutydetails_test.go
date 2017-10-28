package kefudutydetails

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyDetailsRequest(id int32) *KefuDutyDetailsRequest {
	return &KefuDutyDetailsRequest{
		ID: id,
	}
}

func TestGetKefuDutyDetails(t *testing.T) {
	kddr := NewKefuDutyDetailsRequest(326)
	res, err := GetKefuDutyDetails(kddr)
	if err != nil {
		t.Fatalf("TestGetKefuDutyDetails failed:%v", err)
	}
	t.Logf("TestGetKefuDutyDetails return value:%v", res)
}
