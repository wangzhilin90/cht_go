package kefudutydelete

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyDeleteRequest(idstr string) *KefuDutyDeleteRequest {
	return &KefuDutyDeleteRequest{
		Idstr: idstr,
	}
}

func TestDeleteKefuDuty(t *testing.T) {
	kddr := NewKefuDutyDeleteRequest("223,224")
	b := DeleteKefuDuty(kddr)
	if b == false {
		t.Fatalf("TestDeleteKefuDuty failed")
	}
}
