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

func NewMessageUpdateRequestStruct(toUser, oldValue, newValue int32) *MessageUpdateRequestStruct {
	return &MessageUpdateRequestStruct{
		ToUser:        toUser,
		IsPushFlagOld: oldValue,
		IsPushFlagNew: newValue,
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

func TestGetUserDetials(t *testing.T) {
	mirs := NewMessageRequestStruct(475151, "01234567359", "1497542521", 0)
	ms := &messageservice{}
	res, err := ms.GetUserDetials(mirs)
	if err != nil {
		t.Fatalf("TestGetUserInfo failed %v", err)
	}
	t.Log("TestGetUserInfo res %v", res)
}

func TestUpdateMessage(t *testing.T) {
	mur := NewMessageUpdateRequestStruct(9738, 2, 1)
	ms := &messageservice{}
	res, _ := ms.UpdateMessage(mur)
	if res.Status != UPDATE_MES_SUCCESS {
		t.Fatalf("TestUpdateMessage failed")
	}
}
