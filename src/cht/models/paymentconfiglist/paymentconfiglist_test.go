package paymentconfiglist

import (
	_ "cht/initial"
	"testing"
)

func NewPaymentConfigListRequest() *PaymentConfigListRequest {
	return &PaymentConfigListRequest{}
}

func TestGetPaymentConfigList(t *testing.T) {
	pclr := NewPaymentConfigListRequest()
	res, err := GetPaymentConfigList(pclr)
	if err != nil {
		t.Fatalf("TestGetPaymentConfigList failed:%v", err)
	}
	t.Logf("TestGetPaymentConfigList return value:%v", res)
}
