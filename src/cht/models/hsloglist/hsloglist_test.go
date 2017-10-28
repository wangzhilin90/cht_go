package hsloglist

import (
	_ "cht/initial"
	"testing"
)

func NewHsLogListRequest(endtime, utype, limitNum, offset int32) *HsLogListRequest {
	return &HsLogListRequest{
		EndTime:     endtime,
		Utype:       utype,
		LimitNum:    limitNum,
		LimitOffset: offset,
	}
}

func TestGetHsLogTotalNum(t *testing.T) {
	hlr := NewHsLogListRequest(1508482221, 1, 20, 10)
	num, err := GetHsLogTotalNum(hlr)
	if err != nil {
		t.Fatalf("TestGetHsLogTotalNum failed:%v", err)
	}
	t.Logf("TestGetHsLogTotalNum return num:%v", num)
}

func TestGetHsLog(t *testing.T) {
	hlr := NewHsLogListRequest(1508482221, 1, 20, 0)
	res, err := GetHsLog(hlr)
	if err != nil {
		t.Fatalf("TestExportHsLog failed:%v", err)
	}
	t.Logf("TestExportHsLog return value:%v", res)
}
