package hscashlist

import (
	_ "cht/initial"
	"testing"
)

func NewHsCashListRequestStruct(endtime, utype, limitNum, offset int32) *HsCashListRequestStruct {
	return &HsCashListRequestStruct{
		EndTime:     endtime,
		Utype:       utype,
		LimitNum:    limitNum,
		LimitOffset: offset,
	}
}

func TestGetHsCashList(t *testing.T) {
	hlrs := NewHsCashListRequestStruct(1508482221, 1, 20, 10)
	hcls := hscashlistservice{}
	res, _ := hcls.GetHsCashList(hlrs)
	if res.Status != QUERY_HS_CASH_LIST_SUCCESS {
		t.Fatalf("TestGetHsCashList failed")
	}
	t.Logf("TestGetHsCashList response:%v", res)
}
