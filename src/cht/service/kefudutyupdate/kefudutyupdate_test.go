package kefudutyupdate

import (
	_ "cht/initial"
	"testing"
)

func NewKefuDutyUpdateRequestStruct(id int32, customer, holiday_user string, dutytime int32) *KefuDutyUpdateRequestStruct {
	return &KefuDutyUpdateRequestStruct{
		ID:          id,
		Customer:    customer,
		HolidayUser: holiday_user,
		DutyTime:    dutytime,
	}
}

func TestUpdateKefuDuty(t *testing.T) {
	kdurs := NewKefuDutyUpdateRequestStruct(328, "刘刘love,llefr", "xiuxiankefu", 1509347703)
	kdus := kefudutyupdateservice{}
	res, _ := kdus.UpdateKefuDuty(kdurs)
	if res.Status != UPDATE_KEFU_DUTY_SUCCESS {
		t.Fatalf("TestUpdateKefuDuty failed")
	}
	t.Logf("TestUpdateKefuDuty response:%v", res)
}
