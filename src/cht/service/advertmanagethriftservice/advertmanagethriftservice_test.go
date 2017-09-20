package advertmanagethriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertManageRequestStruct() *AdvertManageRequestStruct {
	return &AdvertManageRequestStruct{
		Type:  1,
		Limit: 5,
	}
}

func TestGetAdvertManage(t *testing.T) {
	amrs := NewAdvertManageRequestStruct()
	ams := advertmanageservice{}
	res, err := ams.GetAdvertManage(amrs)
	if err != nil {
		t.Fatalf("TestGetAdvertManage failed")
	}
	t.Logf("TestGetAdvertManage res:%v", res)
}
