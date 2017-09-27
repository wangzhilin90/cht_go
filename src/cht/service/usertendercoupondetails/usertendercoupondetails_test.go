package usertendercoupondetails

import (
	_ "cht/initial"
	"testing"
)

func NewTenderCouponRequestStruct(userid int32, tenderid int32, couponid int32, tendermoney string, timelimit int32) *UserTenderCouponDetailsRequestStruct {
	return &UserTenderCouponDetailsRequestStruct{
		UserId:      userid,
		TenderId:    tenderid,
		CouponId:    couponid,
		TenderMoney: tendermoney,
		TimeLimit:   timelimit,
	}
}

func TestGetRedbagInfo(t *testing.T) {
	tcs := NewTenderCouponRequestStruct(242543, 888, 242533, "20000.35", 19)
	gts := gettendercouponservice{}
	res, err := gts.GetUserTenderCouponDetails(tcs)
	if err != nil {
		t.Fatalf("TestGetRedbagInfo failed:", err)
	}
	t.Log("TestGetRedbagInfo res:", res)
}
