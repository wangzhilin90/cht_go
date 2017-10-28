package hsloglist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	NEXT_DAY_TIME = 24 * 3600
)

type HsLogListRequest struct {
	StartTime            int32
	EndTime              int32
	Type                 int32
	Type2                int32
	Kws                  string
	Utype                int32
	IsExport             int32
	LimitOffset          int32
	LimitNum             int32
	BorrowID             int32
	ChengHuiTongTraceLog string
}

type HsLogDetails struct {
	ID          int32  `orm:"column(id)"`
	UserID      int32  `orm:"column(user_id)"`
	Orderno     string `orm:"column(orderno)"`
	Type        int32  `orm:"column(type)"`
	Money       string `orm:"column(money)"`
	FreezeMoney string `orm:"column(freeze_money)"`
	WaitMoney   string `orm:"column(wait_money)"`
	Addtime     int32  `orm:"column(addtime)"`
	Toid        int32  `orm:"column(toid)"`
	Remark      string `orm:"column(remark)"`
	Username    string `orm:"column(username)"`
	Realname    string `orm:"column(realname)"`
	Regtime     int32  `orm:"column(regtime)"`
}

/**
 * [GetHsLogTotalNum 得到徽商总记录数]
 * @param    hllr *HsLogListRequest 请求入参
 * @return   int32 总记录数
 * @DateTime 2017-10-20T14:20:30+0800
 */
func GetHsLogTotalNum(hllr *HsLogListRequest) (int32, error) {
	Logger.Debugf("GetHsLogTotalNum input param:%v", hllr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	if hllr.StartTime != 0 || hllr.EndTime != 0 || hllr.Type != -1 || (hllr.Type2 != 0 && hllr.Kws != "") || hllr.Utype != 0 {
		Logger.Debugf("GetHsLogTotalNum query condition is not null")
		qb.Select("COUNT(1)").
			From("jl_hs_log HL").
			LeftJoin("jl_user U").On("HL.user_id=U.id").
			Where("1=1")
	} else {
		Logger.Debugf("GetHsLogTotalNum query condition is null")
		qb.Select("COUNT(1)").
			From("jl_hs_log").
			Where("1=1")
	}

	if hllr.StartTime != 0 {
		qb.And(fmt.Sprintf("HL.addtime>=%d", hllr.StartTime))
	}

	if hllr.EndTime != 0 && hllr.StartTime <= hllr.EndTime {
		//需要加一天
		qb.And(fmt.Sprintf("HL.addtime<%d", hllr.EndTime+NEXT_DAY_TIME))
	}

	if hllr.Type != -1 {
		qb.And(fmt.Sprintf("HL.type=%d", hllr.Type))
	}

	if hllr.Type2 != 0 && hllr.Kws != "" {
		if hllr.Type2 == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 2 {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 3 {
			qb.And(fmt.Sprintf("HL.user_id=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 4 {
			qb.And(fmt.Sprintf("HL.orderno=\"%s\"", hllr.Kws))
		}
	}

	if hllr.Utype == 1 {
		qb.And(fmt.Sprintf("U.is_borrower>0"))
	} else if hllr.Utype == 2 {
		qb.And(fmt.Sprintf("U.is_borrower=640"))
	}

	sql := qb.String()
	Logger.Debugf("GetHsLogTotalNum sql:%v", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err != nil {
		Logger.Debugf("GetHelpList query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetHsLogTotalNum return num:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetHsLog 查询徽商日志明细]
 * @param    hllr *HsLogListRequest 请求入参
 * @return   []HsLogDetails 日志明细详情
 * @DateTime 2017-10-20T14:51:34+0800
 */
func GetHsLog(hllr *HsLogListRequest) ([]HsLogDetails, error) {
	Logger.Debugf("ExportHsLog input param:%v", hllr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("HL.*,U.username,U.realname,U.addtime AS regtime FROM jl_hs_log HL LEFT JOIN jl_user U ON HL.user_id=U.id").
		Where("1=1")

	if hllr.StartTime != 0 {
		qb.And(fmt.Sprintf("HL.addtime>=%d", hllr.StartTime))
	}

	if hllr.EndTime != 0 && hllr.StartTime <= hllr.EndTime {
		//需要加一天
		qb.And(fmt.Sprintf("HL.addtime<%d", hllr.EndTime+NEXT_DAY_TIME))
	}

	if hllr.Type != -1 {
		qb.And(fmt.Sprintf("HL.type=%d", hllr.Type))
	}

	if hllr.Type2 != 0 && hllr.Kws != "" {
		if hllr.Type2 == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 2 {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 3 {
			qb.And(fmt.Sprintf("HL.user_id=\"%s\"", hllr.Kws))
		} else if hllr.Type2 == 4 {
			qb.And(fmt.Sprintf("HL.orderno=\"%s\"", hllr.Kws))
		}
	}

	if hllr.Utype == 1 {
		qb.And(fmt.Sprintf("U.is_borrower>0"))
	} else if hllr.Utype == 2 {
		qb.And(fmt.Sprintf("U.is_borrower=0"))
	}

	qb.OrderBy("HL.id").Desc()

	//0:默认不导出,此时limit和offset生效
	if hllr.IsExport == 0 {
		if hllr.LimitNum != 0 {
			qb.Limit(int(hllr.LimitNum))
		}
		if hllr.LimitOffset != 0 {
			qb.Offset(int(hllr.LimitOffset))
		}
	}
	sql := qb.String()
	Logger.Debugf("ExportHsLog sql:%v", sql)
	var hld []HsLogDetails
	_, err := o.Raw(sql).QueryRows(&hld)
	if err != nil {
		Logger.Errorf("ExportHsLog query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("ExportHsLog return value:%v", hld)
	return hld, nil
}
