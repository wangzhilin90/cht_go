package rechargerecordthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewRechargeRecordRequestStruct(user_id, start_time, end_time, query_time, status, offset, limitnum int32, log string) *RechargeRecordRequestStruct {
	return &RechargeRecordRequestStruct{
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

func TestGetRechargeRecordService(t *testing.T) {
	// rrr := NewRechargeRecordRequest(313579141, 1472486400, 1506441600, 0, 0, 0, 0, "")
	rrr := NewRechargeRecordRequestStruct(242972, 1472486400, 1506441600, 0, 0, 2, 10, "")
	rrs := &rechargerecordservice{}
	res, err := rrs.GetRechargeRecord(rrr)
	if err != nil {
		t.Fatal("TestGetRechargeRecordService failed", err)
	}
	t.Log("TestGetRechargeRecordService return value:", res)
}
