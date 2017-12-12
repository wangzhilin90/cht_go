package borrowrepaymentstatistics

import (
	_ "cht/initial"
	"testing"
)

func NewRepaymentStatisticsRequestStruct(user_id int32, status int32) *RepaymentStatisticsRequestStruct {
	return &RepaymentStatisticsRequestStruct{
		UserID: user_id,
		Status: status,
	}
}

func TestGetRepaymentStatisticsDetails(t *testing.T) {
	rsrs := NewRepaymentStatisticsRequestStruct(29, 1)
	brss := borrowrepaymentstatisticsservice{}
	res, _ := brss.GetRepaymentStatisticsList(rsrs)
	if res.Status != QUERY_REPAYMENT_SUCCESS {
		t.Fatalf("TestGetRepaymentStatisticsDetails query failed")
	}
	t.Logf("TestGetRepaymentStatisticsDetails return value:%v", res)
}

func TestGetTotalReplaymentMoney(t *testing.T) {
	rsrs := NewRepaymentStatisticsRequestStruct(29, 1)
	brss := borrowrepaymentstatisticsservice{}
	res, _ := brss.GetTotalReplaymentMoney(rsrs)
	if res.Status != QUERY_TOTAL_REPLAY_SUCCESS {
		t.Fatalf("TestGetTotalReplaymentMoney query failed")
	}
	t.Logf("TestGetTotalReplaymentMoney return value:%v", res)
}
