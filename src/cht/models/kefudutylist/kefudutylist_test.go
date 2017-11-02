package kefudutylist

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyListRequest(kefu int32) *KefuDutyListRequest {
	return &KefuDutyListRequest{
		Kefu: kefu,
	}
}

func TestGetKefuDutyListTotalNum(t *testing.T) {
	kdlr := NewKefuDutyListRequest(15)
	num, err := GetKefuDutyListTotalNum(kdlr)
	if err != nil {
		t.Fatalf("TestGetKefuDutyListTotalNum failed:%v", err)
	}
	t.Logf("TestGetKefuDutyListTotalNum return num:%v", num)
}

func TestGetKefuDutyList(t *testing.T) {
	kdlr := NewKefuDutyListRequest(57)
	res, err := GetKefuDutyList(kdlr)
	if err != nil {
		t.Fatalf("TestGetKefuDutyList failed:%v", err)
	}
	t.Logf("TestGetKefuDutyList return value:%v", res)
}
