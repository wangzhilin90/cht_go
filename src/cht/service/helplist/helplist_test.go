package helplist

import (
	_ "cht/initial"
	"testing"
)

func NewHelpListRequestStruct(status, cateId, limit, offset int32) *HelpListRequestStruct {
	return &HelpListRequestStruct{
		Status:      status,
		Cateid:      cateId,
		LimitNum:    limit,
		LimitOffset: offset,
	}
}

func TestGetHelpList(t *testing.T) {
	hr := NewHelpListRequestStruct(1, 2, 3, 20)
	hs := helplistservice{}
	res, _ := hs.GetHelpList(hr)
	t.Logf("TestGetHelpList return value:%v", res)
}
