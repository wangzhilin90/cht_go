package kefudutyupdate

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyUpdateRequest(id int32, customer, holiday_user string, dutytime int32) *KefuDutyUpdateRequest {
	return &KefuDutyUpdateRequest{
		ID:          id,
		Customer:    customer,
		HolidayUser: holiday_user,
		DutyTime:    dutytime,
	}
}

func TestUpdateKefuDuty(t *testing.T) {
	kdur := NewKefuDutyUpdateRequest(326, "刘刘love,llefr", "xiuxiankefu", 1509347703)
	b := UpdateKefuDuty(kdur)
	if b == false {
		t.Fatalf("TestUpdateKefuDuty failed")
	}
}
