package kefudutyadd

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyAddRequestStruct(customer string) *KefuDutyAddRequestStruct {
	return &KefuDutyAddRequestStruct{
		Customer: customer,
	}
}

func TestAddKefuDuty(t *testing.T) {
	kfdars := NewKefuDutyAddRequestStruct("刘刘love")
	kfdas := kefudutyaddservice{}
	res, _ := kfdas.AddKefuDuty(kfdars)
	if res.Status != KE_FU_DUTY_ADD_SUCCESS {
		t.Fatalf("TestAddKefuDuty failed")
	}
}
