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

func TestGetMessageDetails(t *testing.T) {
	mirs := NewMessageRequestStruct(2332084, "13066008147", "1497542521", 0)
	ms := &messageservice{}
	res, err := ms.GetMessageDetails(mirs)
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

func TestGetUserInfo(t *testing.T) {
	mirs := NewMessageRequestStruct(475151, "01234567359", "1497542521", 0)
	ms := &messageservice{}
	res, err := ms.GetUserInfo(mirs)
	if err != nil {
		t.Fatalf("TestGetUserInfo failed %v", err)
	}
	t.Log("TestGetUserInfo res %v", res)
}
