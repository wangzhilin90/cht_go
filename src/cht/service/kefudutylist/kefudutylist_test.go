package kefudutylist

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyListRequestStruct(kefu int32) *KefuDutyListRequestStruct {
	return &KefuDutyListRequestStruct{
		Kefu: kefu,
	}
}

func TestGetKefuDutyList(t *testing.T) {
	kdlrs := NewKefuDutyListRequestStruct(57)
	kdls := kefudutylistservice{}
	res, _ := kdls.GetKefuDutyList(kdlrs)
	if res.Status != QUERY_KEFU_DUTY_LIST_SUCCESS {
		t.Fatalf("TestGetKefuDutyList failed")
	}
	t.Logf("TestGetKefuDutyList response:%v", res)
}
