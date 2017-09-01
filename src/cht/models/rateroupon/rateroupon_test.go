package rateroupon

import (
	_ "cht/initial"
	"testing"
)

func NewCouponRequest(userId int32, status int32, limit int32, log string, orderby string) *CouponRequest {
	crs := CouponRequest{
		UserID:               userId,
		Status:               status,
		Limit:                limit,
		ChengHuiTongTraceLog: log,
		OrderBy:              orderby,
	}
	return &crs
}

func TestGetRateRoupon(t *testing.T) {
	crs := NewCouponRequest(5004, 3, 10, "", "app_add")
	res, err := GetRateRoupon(crs)
	if err != nil {
		t.Fatalf("GetRateRoupon failed", err)
	}
	for _, v := range res {
		t.Log(v)
	}
}
