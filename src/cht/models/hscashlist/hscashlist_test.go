package hscashlist

import (
	_ "cht/initial"
	"testing"
)

func NewHsCashListRequest(endtime, utype, limitNum, offset int32) *HsCashListRequest {
	return &HsCashListRequest{
		EndTime:     endtime,
		Utype:       utype,
		LimitNum:    limitNum,
		LimitOffset: offset,
	}
}

func TestGetHsCashListTotalNum(t *testing.T) {
	hlr := NewHsCashListRequest(1508482221, 1, 20, 10)
	num, err := GetHsCashListTotalNum(hlr)
	if err != nil {
		t.Fatalf("TestGetHsCashListTotalNum failed:%v", err)
	}
	t.Logf("TestGetHsCashListTotalNum return num:%v", num)
}

func TestGetHsCashList(t *testing.T) {
	hlr := NewHsCashListRequest(1508482221, 1, 20, 10)
	res, err := GetHsCashList(hlr)
	if err != nil {
		t.Fatalf("TestGetHsCashList failed:%v", err)
	}
	t.Logf("TestGetHsCashList return value:%v", res)
}
