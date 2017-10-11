package talking

import (
	_ "cht/initial"
	"testing"
)

func NewTalkingRequestStruct(cateId, status int32) *TalkingRequestStruct {
	return &TalkingRequestStruct{
		Cateid: cateId,
		Status: status,
	}
}

func TestGetTalkingList(t *testing.T) {
	tlrs := NewTalkingRequestStruct(4, 1)
	ts := talkingservcie{}
	res, _ := ts.GetTalkingList(tlrs)
	if res.OneList == nil || res.TalkList == nil {
		t.Fatalf("TestGetTalkingList failed")
	}
	t.Logf("TestGetTalkingList res:%v", res)
}
