package setmsg

import (
	_ "cht/initial"
	"testing"
)

func NewSetMsgDetailsRequest(user_id int32) *SetMsgDetailsRequest {
	return &SetMsgDetailsRequest{
		UserID: user_id,
	}
}

func NewSetMsgDealRequest(user_id int32, status int32, tenderStatus int32) *SetMsgDealRequest {
	return &SetMsgDealRequest{
		UserID:       user_id,
		Status:       status,
		TenderStatus: tenderStatus,
	}
}

func TestGetSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDetailsRequest(1195)
	res, err := GetSetMsgDetails(smdr)
	if err != nil {
		t.Fatalf("TestGetSetMsgDetails failed:%v", err)
	}
	t.Logf("TestGetSetMsgDetails return value:%v", res)
}

func TestUpdateSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDealRequest(20, 3, 2)
	b := UpdateSetMsgDetails(smdr)
	if b == false {
		t.Fatalf("TestUpdateSetMsgDetails failed")
	}
}

func TestInsertSetMsgDetails(t *testing.T) {
	smdr := NewSetMsgDealRequest(20, 1, 3)
	b := InsertSetMsgDetails(smdr)
	if b == false {
		t.Fatalf("TestInsertSetMsgDetails failed")
	}
}
