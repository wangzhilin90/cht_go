package subledgerthriftservice

import (
	. "cht/common/logger"
	// "cht/common/zkclient"
	"cht/models/subledger"
	// "fmt"
	// "git.apache.org/thrift.git/lib/go/thrift"
)

const (
	QUERY_SUBLEDGER_SUCCESS = 1000
	QUERY_SUBLEDGER_FAILED  = 1001
)

var Status = map[int]string{
	QUERY_SUBLEDGER_FAILED:  "查询分账人信息失败",
	QUERY_SUBLEDGER_SUCCESS: "查询分账人信息成功",
}

type subledgerservice struct{}

func (ss *subledgerservice) GetSubledgerList(requestObj *SubledgerRequestStruct) (r *SubledgerResponseStruct, err error) {
	sr := new(subledger.SubledgerRequest)
	sr.HsZhuanrangrenStr = requestObj.GetHsZhuanrangrenStr()
	sr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	sublederList, err := subledger.GetSubledgerList(sr)
	if err != nil {
		Logger.Errorf("GetSubledgerList query failed %v", err)
		return &SubledgerResponseStruct{
			Status: QUERY_SUBLEDGER_FAILED,
			Msg:    Status[QUERY_SUBLEDGER_FAILED],
		}, nil
	}

	var response SubledgerResponseStruct
	for _, v := range sublederList {
		si := new(SubledgerInfoStruct)
		si.UserID = v.UserID
		si.Realname = v.Realname
		si.CardID = v.CardID
		response.SubledgerInfoList = append(response.SubledgerInfoList, si)
	}
	response.Status = QUERY_SUBLEDGER_SUCCESS
	response.Msg = Status[QUERY_SUBLEDGER_SUCCESS]
	Logger.Debugf("GetSubledgerList res:%v", response)
	return &response, nil
}
