package gettenderredbagthriftservice

import (
	. "cht/common/logger"
	// "cht/common/zkclient"
	"cht/models/gettenderredbag"
	// "fmt"
	// "git.apache.org/thrift.git/lib/go/thrift"
)

type gettenderredservice struct{}

const (
	QUERY_RED_BAG_FAILED  = 1001
	QUERY_RED_BAG_SUCCESS = 1002
)

var Status = map[int]string{
	QUERY_RED_BAG_FAILED:  "查询红包金额出错",
	QUERY_RED_BAG_SUCCESS: "查询红包金额成功",
}

func (gts *gettenderredservice) GetRedbagInfo(requestObj *TenderRedbagRequestStruct) (r *TenderRedbagResponseStruct, err error) {
	trr := new(gettenderredbag.TenderRedbagRequest)
	trr.UserId = requestObj.GetUserId()
	trr.TenderId = requestObj.GetTenderId()
	trr.RedId = requestObj.GetRedId()
	trr.TenderMoney = requestObj.GetTenderMoney()
	trr.TimeLimit = requestObj.GetTimeLimit()
	trr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()

	res, err := gettenderredbag.GetRedBagMoney(trr)
	if err != nil {
		return &TenderRedbagResponseStruct{
			Status: QUERY_RED_BAG_FAILED,
			Msg:    Status[QUERY_RED_BAG_FAILED],
		}, nil
	}

	Logger.Debug("GetRedbagInfo red ", res)

	return &TenderRedbagResponseStruct{
		Status:      QUERY_RED_BAG_SUCCESS,
		Msg:         Status[QUERY_RED_BAG_SUCCESS],
		RedbagMoney: res,
	}, nil
}
