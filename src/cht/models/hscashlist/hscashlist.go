package hscashlist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	NEXT_DAY_TIME = 24 * 3600
)

type HsCashListRequest struct {
	StartTime            int32
	EndTime              int32
	Timetype             int32
	Utype                int32 //借款人
	Type                 int32 //类型 用户名,真是姓名，订单号
	Keywords             string
	PayWay               int32 //提现途径
	Status               int32
	IsExport             int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

//徽商提现记录
type HsCashListResult struct {
	ID             int32  `orm:"column(id)"`
	UserID         int32  `orm:"column(user_id)"`
	OrderSn        string `orm:"column(order_sn)"`
	Money          string `orm:"column(money)"`
	Credited       string `orm:"column(credited)"`
	Fee            string `orm:"column(fee)"`
	UseReturnMoney string `orm:"column(use_return_money)"`
	UseFreeNum     int32  `orm:"column(use_free_num)"`
	Addtime        int32  `orm:"column(addtime)"`
	Status         int32  `orm:"column(status)"`
	PayWay         int32  `orm:"column(pay_way)"`
	DealTime       int32  `orm:"column(deal_time)"`
	FailResult_    string `orm:"column(fail_result)"`
	Username       string `orm:"column(username)"`
	Realname       string `orm:"column(realname)"`
	Regtime        int32  `orm:"column(regtime)"`
}

/**
 * [GetHsCashListTotalNum 得到徽商列表记录数]
 * @param     hclr *HsCashListRequest 请求入参
 * @return    int32 返回总数
 * @DateTime 2017-10-23T10:23:41+0800
 */
func GetHsCashListTotalNum(hclr *HsCashListRequest) (int32, error) {
	Logger.Debugf("GetHsCashListTotalNum input param:%v", hclr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(1) FROM jl_hs_cash HC LEFT JOIN jl_user U ON HC.user_id=U.id").
		Where("1=1")

	if hclr.StartTime != 0 {
		if hclr.Timetype == 0 {
			qb.And(fmt.Sprintf("HC.addtime>=%d", hclr.StartTime))
		} else {
			qb.And(fmt.Sprintf("HC.deal_time>=%d", hclr.StartTime))
		}
	}

	if hclr.EndTime != 0 && hclr.StartTime <= hclr.EndTime {
		if hclr.EndTime == 0 {
			qb.And(fmt.Sprintf("HC.addtime<%d", hclr.EndTime+NEXT_DAY_TIME))
		} else {
			qb.And(fmt.Sprintf("HC.deal_time<%d", hclr.EndTime+NEXT_DAY_TIME))
		}
	}

	if hclr.PayWay != -1 {
		qb.And(fmt.Sprintf("HC.pay_way=%d", hclr.PayWay))
	}

	if hclr.Type != 0 && hclr.Keywords != "" {
		if hclr.Type == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", hclr.Keywords))
		} else if hclr.Type == 2 {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", hclr.Keywords))
		} else if hclr.Type == 3 {
			qb.And(fmt.Sprintf("HC.order_sn=\"%s\"", hclr.Keywords))
		}
	}

	if hclr.Utype == 1 {
		//保胜借款人
		qb.And(fmt.Sprintf("U.is_borrower>0"))
	} else if hclr.Utype == 640 {
		//普通客户
		qb.And(fmt.Sprintf("U.is_borrower=640"))
	} else if hclr.Utype == 641 {
		//深圳保胜
		qb.And(fmt.Sprintf("U.is_borrower=641"))
	} else if hclr.Utype == 642 {
		//贵州保胜
		qb.And(fmt.Sprintf("U.is_borrower=642"))
	} else if hclr.Utype == 643 {
		//广州保胜
		qb.And(fmt.Sprintf("U.is_borrower=643"))
	}

	if hclr.Status != -1 {
		qb.And(fmt.Sprintf("HC.status=%d", hclr.Status))
	}

	sql := qb.String()
	Logger.Debugf("GetHsCashListTotalNum sql:%v", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetHsCashListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetHsCashListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetHsCashList 徽商提现记录列表]
 * @param    hclr *HsCashListRequest 请求入参
 * @return   []HsCashListResult 提现记录列表
 * @DateTime 2017-10-23T10:59:04+0800
 */
func GetHsCashList(hclr *HsCashListRequest) ([]HsCashListResult, error) {
	Logger.Debugf("GetHsCashList input param:%v", hclr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("HC.*,U.username,U.realname,U.addtime AS regtime FROM jl_hs_cash HC LEFT JOIN jl_user U ON HC.user_id=U.id").
		Where("1=1")

	if hclr.StartTime != 0 {
		if hclr.Timetype == 0 {
			qb.And(fmt.Sprintf("HC.addtime>=%d", hclr.StartTime))
		} else {
			qb.And(fmt.Sprintf("HC.deal_time>=%d", hclr.StartTime))
		}
	}

	if hclr.EndTime != 0 && hclr.StartTime <= hclr.EndTime {
		if hclr.EndTime == 0 {
			qb.And(fmt.Sprintf("HC.addtime<%d", hclr.EndTime+NEXT_DAY_TIME))
		} else {
			qb.And(fmt.Sprintf("HC.deal_time<%d", hclr.EndTime+NEXT_DAY_TIME))
		}
	}

	if hclr.PayWay != -1 {
		qb.And(fmt.Sprintf("HC.pay_way=%d", hclr.PayWay))
	}

	if hclr.Type != 0 && hclr.Keywords != "" {
		if hclr.Type == 1 {
			qb.And(fmt.Sprintf("U.username=\"%s\"", hclr.Keywords))
		} else if hclr.Type == 2 {
			qb.And(fmt.Sprintf("U.realname=\"%s\"", hclr.Keywords))
		} else if hclr.Type == 3 {
			qb.And(fmt.Sprintf("HC.order_sn=\"%s\"", hclr.Keywords))
		}
	}

	if hclr.Utype == 1 {
		//保胜借款人
		qb.And(fmt.Sprintf("U.is_borrower>0"))
	} else if hclr.Utype == 640 {
		//普通客户
		qb.And(fmt.Sprintf("U.is_borrower=640"))
	} else if hclr.Utype == 641 {
		//深圳保胜
		qb.And(fmt.Sprintf("U.is_borrower=641"))
	} else if hclr.Utype == 642 {
		//贵州保胜
		qb.And(fmt.Sprintf("U.is_borrower=642"))
	} else if hclr.Utype == 643 {
		//广州保胜
		qb.And(fmt.Sprintf("U.is_borrower=643"))
	}

	if hclr.Status != -1 {
		qb.And(fmt.Sprintf("HC.status=%d", hclr.Status))
	}

	qb.OrderBy("HC.id").Desc()

	//0:默认不导出,此时limit和offset生效
	if hclr.IsExport == 0 {
		if hclr.LimitNum != 0 {
			qb.Limit(int(hclr.LimitNum))
		}
		if hclr.LimitOffset != 0 {
			qb.Offset(int(hclr.LimitOffset))
		}
	}
	sql := qb.String()
	Logger.Debugf("GetHsCashList sql:%v", sql)
	var hclres []HsCashListResult
	_, err := o.Raw(sql).QueryRows(&hclres)
	if err != nil {
		Logger.Errorf("GetHsCashList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetHsCashList return value:%v", hclres)
	return hclres, nil
}
