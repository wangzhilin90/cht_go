package goodsedit

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type GoodsEditRequest struct {
	ID                   int32  `thrift:"id,1" db:"id" json:"id"`
	ShowTime             int32  `thrift:"show_time,2" db:"show_time" json:"show_time"`
	CloseTime            int32  `thrift:"close_time,3" db:"close_time" json:"close_time"`
	IsTimer              int32  `thrift:"is_timer,4" db:"is_timer" json:"is_timer"`
	Litpic               string `thrift:"litpic,5" db:"litpic" json:"litpic"`
	Name                 string `thrift:"name,6" db:"name" json:"name"`
	Category             int32  `thrift:"category,7" db:"category" json:"category"`
	RedbagMoney          string `thrift:"redbag_money,8" db:"redbag_money" json:"redbag_money"`
	OriginalPoint        int32  `thrift:"original_point,9" db:"original_point" json:"original_point"`
	CurrentPoint         int32  `thrift:"current_point,10" db:"current_point" json:"current_point"`
	TotalNum             int32  `thrift:"total_num,11" db:"total_num" json:"total_num"`
	SingleNum            int32  `thrift:"single_num,12" db:"single_num" json:"single_num"`
	Content              string `thrift:"content,13" db:"content" json:"content"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,14" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

/**
 * [EditGoods 编辑商品]
 * @param    ger *GoodsEditRequest 请求入参
 * @return   bool true：编辑成功 false:编辑失败
 * @DateTime 2017-10-24T14:35:46+0800
 */
func EditGoods(ger *GoodsEditRequest) bool {
	Logger.Debugf("EditGoods input param:%v", ger)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_point_goods")

	var str string
	if ger.Litpic != "" {
		str += fmt.Sprintf("litpic=\"%s\",", ger.Litpic)
	}

	if ger.Name != "" {
		str += fmt.Sprintf("name=\"%s\",", ger.Name)
	}

	if ger.RedbagMoney != "" {
		str += fmt.Sprintf("redbag_money=\"%s\",", ger.RedbagMoney)
	}

	if ger.Content != "" {
		str += fmt.Sprintf("content=\"%s\",", ger.Content)
	}

	if ger.ShowTime != 0 {
		str += fmt.Sprintf("show_time=%d,", ger.ShowTime)
	}

	if ger.CloseTime != 0 {
		str += fmt.Sprintf("close_time=%d,", ger.CloseTime)
	}

	if ger.IsTimer != 0 {
		str += fmt.Sprintf("is_timer=%d,", ger.IsTimer)
	}

	if ger.Category != 0 {
		str += fmt.Sprintf("category=%d,", ger.Category)
	}

	if ger.OriginalPoint != 0 {
		str += fmt.Sprintf("original_point=%d,", ger.OriginalPoint)
	}

	if ger.CurrentPoint != 0 {
		str += fmt.Sprintf("current_point=%d,", ger.CurrentPoint)
	}

	if ger.TotalNum != 0 {
		str += fmt.Sprintf("total_num=%d,", ger.TotalNum)
	}

	if ger.SingleNum != 0 {
		str += fmt.Sprintf("single_num=%d,", ger.SingleNum)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("id=%d", ger.ID))
	sql := qb.String()

	Logger.Debug("EditGoods sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Errorf("EditGoods update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
