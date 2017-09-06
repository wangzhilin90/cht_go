package gettendercouponthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewTenderCouponRequestStruct(userid int32, tenderid int32, couponid int32, tendermoney string, timelimit int32) *TenderCouponRequestStruct {
	return &TenderCouponRequestStruct{
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
	res, err := gts.GetCouponInfo(tcs)
	if err != nil {
		t.Fatalf("TestGetRedbagInfo failed:", err)
	}
	t.Log("TestGetRedbagInfo res:", res)
}
