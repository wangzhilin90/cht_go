package advertmanage

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertManageRequest() *AdvertManageRequest {
	return &AdvertManageRequest{
		Type:  1,
		Limit: 5,
	}
}

func TestGetAdvertManage(t *testing.T) {
	amr := NewAdvertManageRequest()
	res, err := GetAdvertManage(amr)
	if err != nil {
		t.Fatalf("TestGetAdvertManage failed :%v", err)
	}
	t.Logf("TestGetAdvertManage res:%v", res)
}
