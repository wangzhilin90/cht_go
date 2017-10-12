package helplist

import (
	_ "cht/initial"
	"testing"
)

func NewHelpListRequest(status, cateId int32) *HelpListRequest {
	return &HelpListRequest{
		Status: status,
		Cateid: cateId,
	}
}

func TestGetHelpList(t *testing.T) {
	hr := NewHelpListRequest(1, 2)
	res, err := GetHelpList(hr)
	if err != nil {
		t.Fatalf("TestGetHelpList query failed", err)
	}
	t.Logf("TestGetHelpList return value:%v", res)
}
