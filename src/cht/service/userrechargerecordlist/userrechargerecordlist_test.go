package userrechargerecordlist

import (
	_ "cht/initial"
	"testing"
)

func NewRechargeRecordRequestStruct(user_id, start_time, end_time, query_time, status, offset, limitnum int32, log string) *UserRechargeRecordListRequestStruct {
	return &UserRechargeRecordListRequestStruct{
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

func TestGetUserRechargeRecordList(t *testing.T) {
	// rrr := NewRechargeRecordRequest(313579141, 1472486400, 1506441600, 0, 0, 0, 0, "")
	rrr := NewRechargeRecordRequestStruct(242972, 1472486400, 1506441600, 0, 0, 2, 10, "")
	rrs := &rechargerecordservice{}
	res, err := rrs.GetUserRechargeRecordList(rrr)
	if err != nil {
		t.Fatal("TestgetUserRechargeRecordListService failed", err)
	}
	t.Log("TestgetUserRechargeRecordListService return value:", res)
}
