package operationaldata

import (
	_ "cht/initial"
	"testing"
)

func NewOperationalDataRequestStruct(startMonth, start int32) *OperationalDataRequestStruct {
	return &OperationalDataRequestStruct{
		StartMonth:    startMonth,
		Start:         start,
		TodayTime:     1516614552,
		YesterdayTime: 1316614552,
		TomorrowTime:  0,
	}
}

func TestGetThirtyDaysResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetThirtyDaysResult(odrs)
	if err != nil {
		t.Fatalf("TestGetThirtyDaysResult failed", err)
	}
	t.Logf("TestGetThirtyDaysResult return value:%v", res)
}

func TestGetTwelveMonthResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTwelveMonthResult(odrs)
	if err != nil {
		t.Fatalf("TestGetTwelveMonthResult failed %v", err)
	}
	t.Logf("TestGetTwelveMonthResult return value:%v", res)
}

func TestGetOneMonthResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetOneMonthResult(odrs)
	if err != nil {
		t.Fatalf("TestGetOneMonthResult failed %v", err)
	}
	t.Logf("TestGetOneMonthResult return value:%v", res)
}

func TestGetPeriodResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetPeriodResult(odrs)
	if err != nil {
		t.Fatalf("TestGetPeriodResult failed %v", err)
	}
	t.Logf("TestGetPeriodResult return value:%v", res)
}

func TestGetInvestResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetInvestResult(odrs)
	if err != nil {
		t.Fatalf("TestGetInvestResult failed %v", err)
	}
	t.Logf("TestGetInvestResult return value:%v", res)
}

func TestGetBidResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetBidResult(odrs)
	if err != nil {
		t.Fatalf("TestGetBidResult failed %v", err)
	}
	t.Logf("TestGetBidResult return value:%v", res)
}

func TestGetWaitResult(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetWaitResult(odrs)
	if err != nil {
		t.Fatalf("TestGetWaitResult failed %v", err)
	}
	t.Logf("TestGetWaitResult return value:%v", res)
}

func TestGetTwelveMonthTotalNum(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTwelveMonthTotalNum(odrs)
	if err != nil {
		t.Fatalf("TestGetTwelveMonthTotalNum failed %v", err)
	}
	t.Logf("TestGetTwelveMonthTotalNum return value:%v", res)
}

func TestGetTotalRepayment(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTotalRepayment(odrs)
	if err != nil {
		t.Fatalf("TestGetTotalRepayment failed %v", err)
	}
	t.Logf("TestGetTotalRepayment return value:%v", res)
}

func TestGetTender(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTender(odrs)
	if err != nil {
		t.Fatalf("TestGetTender failed %v", err)
	}
	t.Logf("TestGetTender return value:%v", res)
}

func TestGetTenderToday(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTenderToday(odrs)
	if err != nil {
		t.Fatalf("TestGetTenderToday failed %v", err)
	}
	t.Logf("TestGetTenderToday return value:%v", res)
}

func TestGetProfit(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetProfit(odrs)
	if err != nil {
		t.Fatalf("TestGetProfit failed %v", err)
	}
	t.Logf("TestGetProfit return value:%v", res)
}

func TestGetTenderUserCount(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	res, err := GetTenderUserCount(odrs)
	if err != nil {
		t.Fatalf("TestGetTenderUserCount failed %v", err)
	}
	t.Logf("TestGetTenderUserCount return value:%v", res)
}
