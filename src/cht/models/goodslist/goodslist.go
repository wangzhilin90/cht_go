package goodslist

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type GoodsListRequest struct {
	Name                 string
	Category             int32
	IsExport             int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type GoodsListResult struct {
	ID            int32  `orm:"column(id)"`
	Addtime       int32  `orm:"column(addtime)"`
	ShowTime      int32  `orm:"column(show_time)"`
	CloseTime     int32  `orm:"column(close_time)"`
	IsTimer       int32  `orm:"column(is_timer)"`
	Category      int32  `orm:"column(category)"`
	RedbagMoney   string `orm:"column(redbag_money)"`
	OriginalPoint int32  `orm:"column(original_point)"`
	CurrentPoint  int32  `orm:"column(current_point)"`
	TotalNum      int32  `orm:"column(total_num)"`
	SoldNum       int32  `orm:"column(sold_num)"`
	SingleNum     int32  `orm:"column(single_num)"`
	Name          string `orm:"column(name)"`
	Litpic        string `orm:"column(litpic)"`
	Content       string `orm:"column(content)"`
}

/**
 * [GetGoodsListTotalNum 商品管理列表总记录数]
 * @param    glr *GoodsListRequest 请求入参
 * @return   int32 总数
 * @DateTime 2017-10-24T15:09:38+0800
 */
func GetGoodsListTotalNum(glr *GoodsListRequest) (int32, error) {
	Logger.Debugf("GetGoodsListTotalNum input param:%v", glr)
	o := orm.NewOrm()
	o.Using("default")
	var sql string

	Logger.Debugf("GetGoodsListTotalNum query condition is null")
	buf := bytes.Buffer{}
	buf.WriteString("SELECT COUNT(1) FROM jl_point_goods")
	sql = buf.String()

	Logger.Debugf("GetGoodsListTotalNum sql:%v", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetGoodsListTotalNum query failed :%v", err)
		return 0, err
	}
	Logger.Debugf("GetGoodsListTotalNum return num:%v", totalNum)
	return totalNum, nil
}

func GetGoodsList(glr *GoodsListRequest) ([]GoodsListResult, error) {
	Logger.Debugf("GetGoodsList input param:%v", glr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_point_goods").
		Where("1=1")

	if glr.Name != "" {
		qb.And(fmt.Sprintf("name LIKE \"%%%s%%\"", glr.Name))
	}

	if glr.Category != 0 {
		qb.And(fmt.Sprintf("category=%d", glr.Category))
	}

	qb.OrderBy("id").Desc()

	//0:默认不导出,此时limit和offset生效
	if glr.IsExport == 0 {
		if glr.LimitNum != 0 {
			qb.Limit(int(glr.LimitNum))
		}
		if glr.LimitOffset != 0 {
			qb.Offset(int(glr.LimitOffset))
		}
	}
	sql := qb.String()
	Logger.Debugf("GetGoodsList sql:%v", sql)

	var glrs []GoodsListResult
	_, err := o.Raw(sql).QueryRows(&glrs)
	if err != nil {
		Logger.Errorf("GetGoodsList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetGoodsList return value:%v", glrs)
	return glrs, nil
}
