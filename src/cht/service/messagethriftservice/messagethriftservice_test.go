package messagethriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewMessageRequestStruct(smsid int32, phone string, addtime string, mesType int32) *MessageRequestStruct {
	return &MessageRequestStruct{
		Smsid:   smsid,
		Phone:   phone,
		Addtime: addtime,
		Type:    mesType,
	}
}

func TestGetMessageInfo(t *testing.T) {
	mirs := NewMessageRequestStruct(2332084, "13066008147", "1497542521", 0)
	ms := &messageservice{}
	res, err := ms.GetMessageInfo(mirs)
	if err != nil {
		t.Fatalf("TestGetMessageInfo failed %v", err)
	}
	t.Log("TestGetMessageInfo res %v", res)
}

func TestGetMessageCount(t *testing.T) {
	mirs := NewMessageRequestStruct(2332084, "13066008147", "1497542521", 0)
	ms := &messageservice{}
	res, err := ms.GetMessageCount(mirs)
	if err != nil {
		t.Fatalf("TestGetMessageInfo failed %v", err)
	}
	t.Log("TestGetMessageInfo res %v", res)
}
