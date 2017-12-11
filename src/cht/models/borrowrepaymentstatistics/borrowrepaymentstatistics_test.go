package borrowrepaymentstatistics

import (
	_ "cht/initial"
	"testing"
)

func NewRepaymentStatisticsRequest(user_id int32) *RepaymentStatisticsRequest {
	return &RepaymentStatisticsRequest{
		UserID: user_id,
	}
}

func TestGetRepaymentStatisticsDetails(t *testing.T) {
	rsr := NewRepaymentStatisticsRequest(35)
	res, err := GetRepaymentStatisticsDetails(rsr)
	if err != nil {
		t.Fatalf("TestGetRepaymentStatisticsDetails failed:%v", err)
	}
	t.Log("TestGetRepaymentStatisticsDetails resturn value:", res)
}
