package borrowrepaymentstatistics

import (
	_ "cht/initial"
	"testing"
)

func NewRepaymentStatisticsRequestStruct(user_id int32) *RepaymentStatisticsRequestStruct {
	return &RepaymentStatisticsRequestStruct{
		UserID: user_id,
	}
}

func TestGetRepaymentStatisticsDetails(t *testing.T) {
	rsrs := NewRepaymentStatisticsRequestStruct(29)
	brss := borrowrepaymentstatisticsservice{}
	res, _ := brss.GetRepaymentStatisticsDetails(rsrs)
	if res.Status != QUERY_REPAYMENT_SUCCESS {
		t.Fatalf("TestGetRepaymentStatisticsDetails query failed")
	}
	t.Logf("TestGetRepaymentStatisticsDetails return value:%v", res)
}
