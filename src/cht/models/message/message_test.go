package message

import (
	_ "cht/initial"
	"testing"
)

func NewMessageRequest(smsid int32, phone string, addtime string, mesType int32) *MessageRequest {
	return &MessageRequest{
		Smsid:   smsid,
		Phone:   phone,
		Addtime: addtime,
		Type:    mesType,
	}
}

func TestGetMessageInfo(t *testing.T) {
	mr := NewMessageRequest(2332084, "13066008147", "1497542521", 0)
	res, err := GetMessageInfo(mr)
	if err != nil {
		t.Fatalf("TestGetMessageInfo failed %v", err)
	}
	t.Log("TestGetMessageInfo res %v", res)
}

func TestGetMessageCount(t *testing.T) {
	mr := NewMessageRequest(149754, "13066008147423", "1497542521", 0)
	res, err := GetMessageCount(mr)
	if err != nil {
		t.Fatalf("TestGetMessageCount failed %v", err)
	}
	t.Log("TestGetMessageCount res %v", res)
}
