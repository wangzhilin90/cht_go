package goodsdetails

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type GoodsDetailsRequest struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

type GoodsDetailsStruct struct {
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

func GetGoodsDetails(gdr *GoodsDetailsRequest) (*GoodsDetailsStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetGoodsDetails input param:", gdr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_point_goods").
		Where(fmt.Sprintf("id=%d", gdr.ID))

	sql := qb.String()
	Logger.Debug("GetGoodsDetails sql:", sql)
	var gds GoodsDetailsStruct
	err := o.Raw(sql).QueryRow(&gds)
	if err != nil {
		Logger.Debugf("GetGoodsDetails query failed:", err)
		return nil, err
	}
	Logger.Debugf("GetGoodsDetails res:%v", gds)
	return &gds, nil
}
