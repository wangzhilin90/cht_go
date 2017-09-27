package usertenderredbagdestails

import (
	_ "cht/initial"
	"testing"
)

func NewTenderRedbagRequestStruct(user_id int32, tenderid int32, redid int32, tendermoney string, timelimit int32) *UserTenderRedbagDestailsRequestStruct {
	return &UserTenderRedbagDestailsRequestStruct{
		UserId:      user_id,
		TenderId:    tenderid,
		RedId:       redid,
		TenderMoney: tendermoney,
		TimeLimit:   timelimit,
	}
}

func TestGetUserTenderRedbagDestails(t *testing.T) {
	trs := NewTenderRedbagRequestStruct(242543, 888, 531872, "20000.35", 3)
	gts := gettenderredservice{}
	res, err := gts.GetUserTenderRedbagDestails(trs)
	if err != nil {
		t.Fatalf("TestGetRedbagInfo failed:", err)
	}
	t.Log("TestGetRedbagInfo res:", res)
}
