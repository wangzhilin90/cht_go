package cashrecordthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewCashRecordRequestStruct(user_id, start_time, end_time, query_time, status, offset, limitnum int32, log string) *CashRecordRequestStruct {
	return &CashRecordRequestStruct{
		UserID:               user_id,
		StartTime:            start_time,
		EndTime:              end_time,
		QueryTime:            query_time,
		RechargeStatus:       status,
		LimitOffset:          offset,
		LimitNum:             limitnum,
		ChengHuiTongTraceLog: log,
	}
}

func TestGetCashRecord(t *testing.T) {
	rrr := NewCashRecordRequestStruct(221889, 1376279359, 1376299359, 0, 0, 1, 1, "")
	rrs := &cashrecordservice{}
	res, err := rrs.GetCashRecord(rrr)
	if err != nil {
		t.Fatal("TestGetCashRecord failed:", err)
	}
	t.Log("TestGetCashRecord return value:", res)
}
