package couponlistthriftservice

import (
	_ "cht/initial"
	"testing"
)

func TestGetCoupon(t *testing.T) {
	crs := NewCouponRequestStruct(1, 3, 10, "", "app_add")
	cs := &CouponService{}
	res, err := cs.GetCoupon(crs)
	if err != nil {
		t.Fatalf("GetCoupon failed", err)
	}
	for _, v := range res.CouponList {
		t.Log(v)
	}
}

func TestStartCouponServer(t *testing.T) {
	StartCouponServer()
}
