package talking

import (
	_ "cht/initial"
	"testing"
)

func NewTalkingRequest(cateId, status int32) *TalkingRequest {
	return &TalkingRequest{
		Cateid: cateId,
		Status: status,
	}
}

func TestGetTalkinglist(t *testing.T) {
	tr := NewTalkingRequest(4, 1)
	res, err := GetTalkinglist(tr)
	if err != nil || res == nil {
		t.Fatalf("TestGetTalkinglist failed")
	}
	t.Logf("TestGetTalkinglist res:%v", res)
}

func TestGetOnelist(t *testing.T) {
	tr := NewTalkingRequest(641, 1)
	res, err := GetOnelist(tr)
	if err != nil || res == nil {
		t.Fatalf("TestGetOnelist failed")
	}
	t.Logf("TestGetOnelist res:%v", res)
}
