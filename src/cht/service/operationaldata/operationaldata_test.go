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

func TestGetOperationalData(t *testing.T) {
	odrs := NewOperationalDataRequestStruct(1476176400, 1505120400)
	ods := operationaldataservice{}
	res, _ := ods.GetOperationalData(odrs)
	t.Logf("TestGetOperationalData return value:%v", res)
}
