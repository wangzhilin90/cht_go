package vipcustomerloglist

import (
	_ "cht/initial"
	"testing"
)

func NewVipCustomerLogListRequestStruct(uType int32, keywords string, limitnum, limitOffset int32) *VipCustomerLogListRequestStruct {
	return &VipCustomerLogListRequestStruct{
		Type:        uType,
		Keywords:    keywords,
		LimitNum:    limitnum,
		LimitOffset: limitOffset,
	}
}

func TestGetVipCustomerLogList(t *testing.T) {
	vclr := NewVipCustomerLogListRequestStruct(1, "zengjie001", 20, 0)
	vcls := vipcustomerloglistservice{}
	res, _ := vcls.GetVipCustomerLogList(vclr)
	if res.Status != QUERY_VIP_CUSTOMER_LOG_LIST_SUCCESS {
		t.Fatalf("TestGetVipCustomerLogList failed")
	}
	t.Logf("TestGetVipCustomerLogList return value:%v", res)
}
