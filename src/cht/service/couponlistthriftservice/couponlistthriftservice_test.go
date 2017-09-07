package couponlistthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewCouponRequestStruct(userId int32, status int32, limit int32, log string, orderby string) *CouponRequestStruct {
	crs := CouponRequestStruct{
		UserID:               userId,
		Status:               status,
		Limit:                limit,
		ChengHuiTongTraceLog: log,
		OrderBy:              orderby,
	}
	return &crs
}

func TestGetCoupon(t *testing.T) {
	// crs := NewCouponRequestStruct(1, 3, 10, "", "app_add")
	crs := NewCouponRequestStruct(28, 0, 0, "", "end_time")
	cs := &CouponService{}
	res, err := cs.GetCoupon(crs)
	if err != nil {
		t.Fatalf("GetCoupon failed", err)
	}
	for _, v := range res.CouponList {
		t.Log(v)
	}
}
