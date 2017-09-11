package cashrecord

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

func TestGetCashStats(t *testing.T) {
	crrs := NewCashRecordRequestStruct(2218891, 1472486400, 1506441600, 0, 0, 2, 10, "")
	res, err := GetCashStats(crrs)
	if err != nil {
		t.Fatalf("TestGetCashStats failed", err)
	}
	t.Log("TestGetCashStats return value:", res)
}

func TestGetCashRecord(t *testing.T) {
	// rrr := NewRechargeRecordRequest(313579141, 1472486400, 1506441600, 0, 0, 0, 0, "")
	crrs := NewCashRecordRequestStruct(221889, 1472486400, 1506441600, 0, 0, 2, 10, "")
	res, cs, num, err := GetCashRecord(crrs)
	if err != nil {
		t.Fatal("TestGetCashRecord failed", err)
	}
	t.Log("TestGetCashRecord return value:", res, cs, num)
}
