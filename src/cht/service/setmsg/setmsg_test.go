package setmsg

import (
	_ "cht/initial"
	"testing"
)

func NewSetMsgDetailsRequestStruct(user_id int32) *SetMsgDetailsRequestStruct {
	return &SetMsgDetailsRequestStruct{
		UserID: user_id,
	}
}

func NewSetMsgDealRequestStruct(user_id int32, status int32, tenderStatus int32) *SetMsgDealRequestStruct {
	return &SetMsgDealRequestStruct{
		UserID:       user_id,
		Status:       status,
		TenderStatus: tenderStatus,
	}
}

func TestGetSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDetailsRequestStruct(11)
	sms := setmsgservice{}
	res, _ := sms.GetSetMsgDetails(smdr)
	if res.Status != QUERY_SET_MSG_DETAILS_SUCCESS {
		t.Fatalf("TestGetSetMsgDetails failed")
	}
	t.Logf("TestGetSetMsgDetails return value:%v", res)
}

func TestUpdateSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDealRequestStruct(20, 3, 2)
	sms := setmsgservice{}
	res, _ := sms.UpdateSetMsgDetails(smdr)
	if res.Status != UPDATE_SET_MSG_SUCCESS {
		t.Fatalf("TestUpdateSetMsgDetails failed")
	}
}

func TestInsertSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDealRequestStruct(21, 1, 3)
	sms := setmsgservice{}
	res, _ := sms.InsertSetMsgDetails(smdr)
	if res.Status != INSERT_SET_MSG_SUCCESS {
		t.Fatalf("TestInsertSetMsgDetails failed")
	}
}
