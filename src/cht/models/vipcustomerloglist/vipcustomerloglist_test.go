package vipcustomerloglist

import (
	_ "cht/initial"
	"testing"
)

func NewVipCustomerLogListRequest(uType int32, keywords string, limitnum, limitOffset int32) *VipCustomerLogListRequest {
	return &VipCustomerLogListRequest{
		Type:        uType,
		Keywords:    keywords,
		LimitNum:    limitnum,
		LimitOffset: limitOffset,
	}
}

func TestGetVipCustomerLogListTotalNum(t *testing.T) {
	vclr := NewVipCustomerLogListRequest(1, "zengjie001", 20, 10)
	num, err := GetVipCustomerLogListTotalNum(vclr)
	if err != nil {
		t.Fatalf("TestGetVipCustomerLogListTotalNum failed:%v", err)
	}
	t.Logf("TestGetVipCustomerLogListTotalNum return num:%v", num)
}

func TestGetVipCustomerLogList(t *testing.T) {
	vclr := NewVipCustomerLogListRequest(1, "zengjie001", 20, 0)
	num, err := GetVipCustomerLogList(vclr)
	if err != nil {
		t.Fatalf("TestGetVipCustomerLogList failed:%v", err)
	}
	t.Logf("TestGetVipCustomerLogList return value:%v", num)
}
