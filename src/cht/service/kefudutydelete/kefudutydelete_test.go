package kefudutydelete

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyDeleteRequestStruct(idStr string) *KefuDutyDeleteRequestStruct {
	return &KefuDutyDeleteRequestStruct{
		Idstr: idStr,
	}
}

func TestDeleteKefuDuty(t *testing.T) {
	kddrs := NewKefuDutyDeleteRequestStruct("221,222")
	kfdds := kefudutydeleteservice{}
	res, _ := kfdds.DeleteKefuDuty(kddrs)
	if res.Status != KEFU_DUTY_DELETE_SUCCESS {
		t.Fatalf("TestDeleteKefuDuty failed")
	}
	t.Logf("TestDeleteKefuDuty response:%v", res)
}
