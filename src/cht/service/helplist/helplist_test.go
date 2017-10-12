package helplist

import (
	_ "cht/initial"
	"testing"
)

func NewHelpListRequestStruct(status, cateId int32) *HelpListRequestStruct {
	return &HelpListRequestStruct{
		Status: status,
		Cateid: cateId,
	}
}

func TestGetHelpList(t *testing.T) {
	hr := NewHelpListRequestStruct(1, 2)
	hs := helplistservice{}
	res, _ := hs.GetHelpList(hr)
	t.Logf("TestGetHelpList return value:%v", res)
}
