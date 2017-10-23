package hsloglist

import (
	_ "cht/initial"
	"testing"
)

func NewHsLogListRequestStruct(endtime, utype, limitNum, offset int32) *HsLogListRequestStruct {
	return &HsLogListRequestStruct{
		EndTime:     endtime,
		Utype:       utype,
		LimitNum:    limitNum,
		LimitOffset: offset,
	}
}

func TestGetHslogList(t *testing.T) {
	hlr := NewHsLogListRequestStruct(1508482221, 1, 20, 10)
	hlls := hsloglistservice{}
	res, _ := hlls.GetHslogList(hlr)
	if res.Status != QUERY_HS_LOG_SUCCESS {
		t.Fatalf("TestGetHslogList failed")
	}
	t.Logf("TestGetHslogList return value:%v", res)
}
