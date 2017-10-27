package kefudutyadd

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyAddRequest(customer string) *KefuDutyAddRequest {
	return &KefuDutyAddRequest{
		Customer:    customer,
		HolidayUser: " ",
	}
}

func TestAddKefuDuty(t *testing.T) {
	kdar := NewKefuDutyAddRequest("刘刘")
	b := AddKefuDuty(kdar)
	if b == false {
		t.Fatalf("TestAddKefuDuty failed")
	}
}
