package rechargerecord

import (
	_ "cht/initial"
	"testing"
)

func NewRechargeRecordRequest(user_id, start_time, end_time, query_time, status, offset, limitnum int32, log string) *RechargeRecordRequest {
	return &RechargeRecordRequest{
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

func TestGetRechargeRecord(t *testing.T) {
	// rrr := NewRechargeRecordRequest(313579141, 1472486400, 1506441600, 0, 0, 0, 0, "")
	rrr := NewRechargeRecordRequest(242972, 1472486400, 1506441600, 0, 0, 2, 10, "")
	res, num, err := GetRechargeRecord(rrr)
	if err != nil {
		t.Fatal("TestGetRechargeRecord failed", err)
	}
	t.Log("TestGetRechargeRecord return value:", res, num)
}
