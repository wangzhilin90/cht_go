package borrowrepaymentstatistics

import (
	_ "cht/initial"
	"testing"
)

func NewRepaymentStatisticsRequest(user_id int32, status int32) *RepaymentStatisticsRequest {
	return &RepaymentStatisticsRequest{
		UserID: user_id,
		Status: status,
	}
}

func TestGetRepaymentStatisticsDetails(t *testing.T) {
	rsr := NewRepaymentStatisticsRequest(35, 0)
	res, err := GetRepaymentStatisticsList(rsr)
	if err != nil {
		t.Fatalf("TestGetRepaymentStatisticsDetails failed:%v", err)
	}
	t.Log("TestGetRepaymentStatisticsDetails return value:", res)
}

func TestGetTotalReplaymentMoney(t *testing.T) {
	rsr := NewRepaymentStatisticsRequest(35, 1)
	res, err := GetTotalReplaymentMoney(rsr)
	if err != nil {
		t.Fatalf("TestGetTotalReplaymentMoney failed:%v", err)
	}
	t.Logf("TestGetTotalReplaymentMoney return value:%v", res)
}
