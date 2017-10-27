package paymentconfiglist

import (
	_ "cht/initial"
	"testing"
)

func NewPaymentConfigListRequestStruct() *PaymentConfigListRequestStruct {
	return &PaymentConfigListRequestStruct{}
}

func TestGetPaymentConfigList(t *testing.T) {
	pclrs := NewPaymentConfigListRequestStruct()
	pcls := paymentconfiglistservice{}
	res, _ := pcls.GetPaymentConfigList(pclrs)
	if res.Status != QUERY_PAYMENT_CONFIG_LIST_SUCCESS {
		t.Fatalf("TestGetPaymentConfigList failed")
	}
	t.Logf("TestGetPaymentConfigList response:%v", res)
}
