package gettenderredbag

import (
	_ "cht/initial"
	"testing"
)

func NewTenderRedbagRequest(user_id int32, tenderid int32, redid int32, tendermoney string, timelimit int32) *TenderRedbagRequest {
	return &TenderRedbagRequest{
		UserId:      user_id,
		TenderId:    tenderid,
		RedId:       redid,
		TenderMoney: tendermoney,
		TimeLimit:   timelimit,
	}
}

func TestGetBorrowType(t *testing.T) {
	trr := NewTenderRedbagRequest(242543, 888, 531872, "20000.35", 3)
	res, err := getBorrowType(trr)
	if err != nil {
		t.Fatalf("TestGetBorrowType failed", err)
	}
	t.Log("TestGetBorrowType res ", res)
}

func TestGetRedBagMoney(t *testing.T) {
	trr := NewTenderRedbagRequest(242543, 888, 531872, "20000.35", 3)
	res, err := GetRedBagMoney(trr)
	if err != nil {
		t.Fatalf("TestGetRedBagMoney failed ", err)

	}
	t.Log("TestGetRedBagMoney res ", res)
}
