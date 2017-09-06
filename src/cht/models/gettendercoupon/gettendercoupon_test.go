package gettendercoupon

import (
	_ "cht/initial"
	"testing"
)

func NewTenderCouponRequest(user_id int32, tenderid int32, coupon_id int32, tendermoney string, timelimit int32) *TenderCouponRequest {
	return &TenderCouponRequest{
		UserId:      user_id,
		TenderId:    tenderid,
		CouponId:    coupon_id,
		TenderMoney: tendermoney,
		TimeLimit:   timelimit,
	}
}

func TestGetBorrowType(t *testing.T) {
	trr := NewTenderCouponRequest(242543, 888, 242533, "20000.35", 19)
	res, err := getBorrowType(trr)
	if err != nil {
		t.Fatalf("TestGetBorrowType failed", err)
	}
	t.Log("TestGetBorrowType res ", res)
}

func TestGetTenderCoupon(t *testing.T) {
	tcr := NewTenderCouponRequest(242543, 888, 242533, "20000.35", 19)
	res, err := GetTenderCoupon(tcr)
	if err != nil {
		t.Fatalf("TestGetTenderCoupon failed ", err)

	}
	t.Log("TestGetTenderCoupon res ", res)
}
